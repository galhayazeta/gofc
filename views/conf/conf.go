package conf

import (
	"log"
	"os"

	"github.com/go-ini/ini"
)

var IniConf *ini.File
func init() {
	cfg_, err := ini.Load("conf/conf.ini")
	if err != nil {
		log.Printf("config.ini读取错误: %v", err)
		os.Exit(1)
	}	
	IniConf = cfg_
}