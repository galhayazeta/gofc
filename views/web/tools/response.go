package tools

import "github.com/gin-gonic/gin"


func Response(code int, msg string, data any) (int,any){
	result :=gin.H{
		"code":code,
		"msg":msg,
		"data":data,
	}
	return 200, result
}