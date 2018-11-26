package main

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "os"
	"log"
)

// type Recurlyservers struct {
//     XMLName     xml.Name `xml:"servers"`
//     Version     string   `xml:"version,attr"`
//     Svs         []server `xml:"server"`
//     Description string   `xml:",innerxml"`
// }
//
// type server struct {
//     XMLName    xml.Name `xml:"server"`
//     ServerName string   `xml:"serverName"`
//     ServerIP   string   `xml:"serverIP"`
// }


type xmlResult struct {
	XMLName   xml.Name  `xml:"xmlResult"`
	Handup    handupData `xml:"data"`
}

type handupData struct {
	XMLName   xml.Name  `xml:"data"`
	Stus     []stuhandup `xml:"studentHandUpStatus"`
}

type stuhandup struct {
	XMLName xml.Name `xml:"studentHandUpStatus"`
	ApId   string   `xml:"apId"`
	Classname string  `xml:"classname"`
	EnableHandup  string `xml:"enableHandUp"`
	ID        string   `xml:"id"`
	Name      string   `xml:"name"`
	Seat      string   `xml:"seat"`
}

func main() {
    file, err := os.Open("servers.xml") // For read access.
    if err != nil {
        fmt.Printf("error: %v", err)
        return
    }
    defer file.Close()
    data, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Printf("error: %v", err)
        return
    }
    v := xmlResult{}
    err = xml.Unmarshal(data, &v)
    if err != nil {
        fmt.Printf("error: %v", err)
        return
    }
    fmt.Println(v)
	for i, val := range v.Handup.Stus {
		log.Println(val.ID)
		if val.ID == "00:1a:4a:16:01:6e"  &&
		   val.EnableHandup == "true" {
		   log.Println("xxxxxxxxxxxxxxxxxxx")
		   v.Handup.Stus[i].EnableHandup = "false"
	   }
		tmp := val.ID;
		tmp += " ---- "
		tmp += val.EnableHandup
		log.Println(tmp)
	}
	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(v, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(xml.Header)
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		//写入文件
		ioutil.WriteFile("servers.xml", xmlOutPutData, os.ModeAppend)

		fmt.Println("OK~")
	} else {
		fmt.Println(outPutErr)
	}
}
