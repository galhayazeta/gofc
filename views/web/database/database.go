package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
var DB *gorm.DB

func init(){
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}
	
	db.AutoMigrate(&User_Datas_models{},&Up_Datas_models{},&Video_Datas_models{},&Live_Conf_Datas_models{})
	DB = db

	fc2_default()
	up_init()
}


func fc2_default(){
	p := Live_Conf_Datas_models{}
	d := DB.Where("website = ?", "fc2").First(&p)
	if d.RowsAffected == 0{
		p = Live_Conf_Datas_models{
			Proxy: "",
			Quality: "2Mbps",
			Website:"fc2",
		}
		d2 := DB.Create(&p)
		if d2.RowsAffected == 0{
			log.Fatal("创建fc2配置失败")
		}
	}

}


func up_init(){
	L:=[]Up_Datas_models{}
	DB.Where("state = ?","录播中").Find(&L)
	for _,v := range L{
		v.State = "监控中"
		DB.Save(&v)
	}
}