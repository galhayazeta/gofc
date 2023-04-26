package user

import (
	"api/views/logs"
	"api/views/web/database"
	"api/views/web/tools"
	"fmt"

	"github.com/gin-gonic/gin"
)

type User_func_class struct {
	

}

var User_Api = User_func_class{}

func (obj *User_func_class) Login_POST(req *gin.Context) {
	//用户登录
	json_data := Login_post_json{}
	if err := req.ShouldBindJSON(&json_data); err != nil {
		logs.Log_write("登录", "Login_POST", "用户提交数据格式错误")
		req.JSON(tools.Response(503, "提交格式错误", ""))
		return
	}
	if json_data.Username == "" {
		req.JSON(tools.Response(503, "提交格式错误", gin.H{"username_err": "请输入用户名"}))
		return
	}
	if json_data.Password == "" {
		req.JSON(tools.Response(503, "提交格式错误", gin.H{"password_err": "请输入密码"}))
		return
	}

	userModel := database.User_Datas_models{}
	d := database.DB.Where("username == ?", json_data.Username).First(&userModel)
	if d.RowsAffected == 0 {
		req.JSON(tools.Response(503, "用户不存在", gin.H{"username_err": "用户不存在"}))
		return
	}

	if userModel.Password != json_data.Password {
		logs.Log_write("登录", "Login_POST", fmt.Sprintf("用户%v输入登录密码错误", json_data.Username))
		req.JSON(tools.Response(503, "密码错误", gin.H{"password_err": "密码错误"}))
		return
	}

	s, _ := tools.GenToken(userModel.Username, userModel.Level)

	req.JSON(tools.Response(200, "登录成功", gin.H{
		"token": s,
		"level": userModel.Level,
		"state": userModel.State,
		"username": userModel.Username,
	}))
	return
}

