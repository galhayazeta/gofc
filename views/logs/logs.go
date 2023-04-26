package logs

import (
	"api/views/web/tools"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var Log_char = make(chan string, 20000)


func Log_init() {
	gin.DisableConsoleColor()
	f, err := os.OpenFile("./logs/gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("创建gin日志失败 原因为:%v", err))
	}
	gin.DefaultWriter = io.MultiWriter(f)

}

func Log_write(log_type string, func_name string, msg string) {
	now_str := time.Now().In(tools.Time_zone).Format("2006-01-02 15:04:05")
	data := fmt.Sprintf("%v----%v----%v----%v\n", now_str, log_type, func_name, msg)
	log.Println(data)
	Log_char <- data
}

func Log_write_file() {

	var file *os.File
	gsh := "2006-01" 

	log_file_name := "./logs/" + time.Now().In(tools.Time_zone).Format(gsh) + ".log"
	file, err := os.OpenFile(log_file_name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Println("日志打开失败 原因为:%v", err)
		return 
	}
	for {
		data := <-Log_char
		now_log_file_name := "./logs/" + time.Now().In(tools.Time_zone).Format(gsh) + ".log"
		if log_file_name != now_log_file_name {
	
			file.Close()
			log_file_name = now_log_file_name
			file, err = os.OpenFile(log_file_name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
			if err != nil {
				log.Println("日志打开失败 原因为:%v", err)
				continue
			}
		}
		_, err2 := file.Write([]byte(data))
		if err2 != nil {
			log.Println("日志写入失败%v",err2)
		}
	}
}
