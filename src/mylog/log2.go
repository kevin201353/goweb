package mylog

import (
	"fmt"
	"log"
	"os"
)

const (
	LOG_FILE_PATH = "my_log"
)

func Log2(a ...interface{}) {
	filename := LOG_FILE_PATH
	logfile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0)
	defer logfile.Close()
	if err != nil {
		log.Fatalln("open file error ! \n")
	}
	debuglog := log.New(logfile, "[Debug]", log.Ldate|log.Ltime|log.Llongfile)
	debuglog.Println(a)
	fmt.Println(a)
}
