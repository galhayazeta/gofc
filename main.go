package main

import (
	"api/views/app/fc2"
	"api/views/conf"
	"api/views/logs"
	"api/views/web/route"
	"log"
)



func main(){
	log.Println(conf.IniConf.Section("server").Key("test").String())
	go logs.Log_write_file()
	go fc2.Fc2_main()
	route.Route()
}