package admin

import (
	"api/views/logs"
	"api/views/web/database"
	"fmt"
	"strconv"
	"time"
)

type Admin_User_list struct {
	
}

var Admin_User_api = Admin_User_list{}

func (obj *Admin_User_list)Admin_Create(){
	
	p:= database.User_Datas_models{}
	d := database.DB.Where("level = ?", 1).First(&p)
	if d.RowsAffected == 0{
		timestamp := int32(time.Now().Unix())
		pwd := strconv.Itoa(int(timestamp))
		submit := database.User_Datas_models{
			Username: "admin",
			Password: pwd,
			Level: 1,
		}
		database.DB.Create(&submit)
		logs.Log_write("创建管理员账号", "admin_Create", fmt.Sprintf("管理员密码为:%v", pwd))
	}
}



