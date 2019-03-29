package main

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"encoding/base64"
	_ "go-sqlite3"
	"html/template"
	"mylog"
	"os"
	"reflect"
	"unsafe"
	"strconv"
	"time"
)

const maxUploadSize = 2 * 1024 * 2014 // 2 MB
const uploadPath = "./tmp"

type LoginInfo struct {
	XMLName   xml.Name `xml."Login"`
	Cmd       int 
	Operator  string
	Sessionid string
	Logtype   string
	Tel       string
}

type Photo struct {
	Jpg  string
	Sign string
	Cam  int
}

var db *sql.DB
var err error

func stringtoslicebyte(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func slicebytetostring(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{
		Data: bh.Data,
		Len:  bh.Len,
	}
	return *(*string)(unsafe.Pointer(&sh))
}

func DealRequst(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	login2 := LoginInfo{}
	if r.Method == "GET" {
		login := LoginInfo{}
		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
			login.Operator = r.Form.Get("operatorno")
			login.Logtype = r.Form.Get("logtype")
			login.Sessionid = r.Form.Get("sessionid")
			login.Tel = r.Form.Get("tel")
			cmd, _ := strconv.Atoi(r.Form.Get("cmd"))
			login.Cmd = cmd 
		}
		output, err := xml.MarshalIndent(login, "  ", "    ")
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		err = ioutil.WriteFile("LoginInfo.xml", output, 0777)
		if err != nil {
			fmt.Printf("write file error: %v\n", err)
		}
		login2 = login
	}
	fmt.Fprintf(w, "Hello astaxie!")
	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	if err != nil {
		log.Fatal(err)
	}

	var bExist bool

    for rows.Next() {
		var uid int
		var username string
		var sessionid string
		var logtype string
		var created string
		var tel string
		err = rows.Scan(&uid, &username, &sessionid, &logtype, &tel, &created)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(sessionid)
		fmt.Println(tel)
		fmt.Println(created)


        fmt.Println("login2: ", login2.Operator)
		//检测用户名是否存在，存在且sessionid也存在登录成功
    	if login2.Operator == username && login2.Sessionid == sessionid {
            //成功
            bExist = true
            fmt.Sprintf("%s 登录成功", username)
    	}
    }

    if !bExist { 
         fmt.Println("新增记录")
        stmt, err := db.Prepare("insert into userinfo(user, sessionid, logtype, tel, created) values(?,?,?,?,?)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		res, err := stmt.Exec(login2.Operator, login2.Sessionid, login2.Logtype, login2.Tel, time.Now())
		if err != nil {
			log.Fatal(err)
		}
		id, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(id)
    }
}


func OnInput(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		fmt.Println(r.Form)
		fmt.Println("path", r.URL.Path)
		fmt.Println("scheme", r.URL.Scheme)
		tmpl := template.Must(template.ParseFiles("input.html"))
		err := tmpl.Execute(w, nil)
		if err != nil {
			log.Fatalf("template execution: %s", err)
		}
	} else {
		fmt.Println("post")
		r.ParseMultipartForm(maxUploadSize)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println("form file err:", err)
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		fmt.Println(handler.Header)
		f, err := os.OpenFile("./files/"+handler.Filename, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("open file err:", err)
			return
		}
		defer f.Close()
		bytes, err := io.Copy(f, file)
		if err != nil {
			log.Fatal("io copy err: ", err)
		}
		fmt.Println("io copy file bytes: ", bytes)
	}
}

func uploadjson(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		fmt.Println(r.Form)
	} else {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("uploadPhoto Read failed:", err)
		}
		defer r.Body.Close()
		//fmt.Println("json data: ", b)
		mylog.Log2("json data: ", b)
		var photo Photo
		err = json.Unmarshal(b, &photo)
		if err != nil {
			log.Fatal("json unmarshal err:", err)
			return
		}
		fmt.Println("photo Jpg:", photo.Jpg)
		mylog.Log2("photo Jpg: ", photo.Jpg, "Sign ", photo.Sign, "Cam: ", photo.Cam)

		/*
			decoded, err2 := base64.StdEncoding.DecodeString(photo.Jpg)
			if err != nil {
				fmt.Println("decode error:", err)
				return
			}
			err2 = ioutil.WriteFile("./files/photo2.jpg", decoded, 0666)
			if err2 != nil {
				log.Fatal("write photo file err:", err2)
				return
			}*/
	}
}

func upgernal(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		fmt.Println(r.Form)
	} else {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("upgernal Read failed:", err)
		}
		defer r.Body.Close()
		data := slicebytetostring(b)
		//fmt.Println("data 111: ", data)
		mylog.Log2("data 111: ", data)
		decoded, err2 := base64.StdEncoding.DecodeString(data)
		if err != nil {
			fmt.Println("decode error:", err)
			return
		}
		err2 = ioutil.WriteFile("./files/photo.jpg", decoded, 0666)
		if err2 != nil {
			log.Fatal("write photo file err:", err2)
			return
		}
	}
}

func getAllinfor(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
	} else {
		//fmt.Println("getAllinfor")
		//w.Write(stringtoslicebyte("getAllinfor"))
		for {
			//time.Sleep(3000 * time.Millisecond)
		}
	}
}

func main() {
	db, err = sql.Open("sqlite3", "my.db")
	if err != nil {
		log.Fatal("sql open: ", err)
	}
	//创建表
	sql_table := `
    CREATE TABLE IF NOT EXISTS userinfo(
        uid INTEGER PRIMARY KEY AUTOINCREMENT,
        user VARCHAR(64) NULL,
        sessionid VARCHAR(64) NULL,
		logtype  VARCHAR(64) NULL,
        tel  VARCHAR(64) NULL,
        created  DATE  NULL
    );
    `
	_, err = db.Exec(sql_table)
	if err != nil {
		log.Printf("%q: %s\n", err, sql_table)
		return
	}
	http.HandleFunc("/", DealRequst)
	//http.HandleFunc("/", OnInput)
	http.HandleFunc("/uploadJson", uploadjson)
	http.HandleFunc("/upgernal", upgernal)
	http.HandleFunc("/getAllinfor", getAllinfor)
	err = http.ListenAndServe(":8008", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
