package admin

import (
	"api/views/web/tools"
	"strings"

	"github.com/gin-gonic/gin"
)

func Handler_Admin_auth() gin.HandlerFunc {
	
    return func(ctx *gin.Context) {
		dic, err := tools.ParseToekn(ctx.GetHeader("Authorization"))


		
		if strings.Contains(ctx.Request.URL.String(), "/video/down"){
			

		}else{
			if err != nil{
				ctx.Abort()   
				ctx.JSON(tools.Response(403, "请重新登录", ""))
				return
			}
			if dic.Level != 1{
				ctx.Abort() 
				ctx.JSON(tools.Response(403, "请重新登录", ""))
				return
			}
		}



        
        ctx.Next()
        
		
    }
}





