package route

import (
	"api/views/conf"
	"api/views/logs"
	"api/views/web/pages/admin"
	"api/views/web/pages/live"
	"api/views/web/pages/liveset"
	"api/views/web/pages/user"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os/exec"
	"runtime"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)



func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}




func Route() {
	logs.Log_init()
	r := gin.Default()
	_ = mime.AddExtensionType(".js", "application/javascript")


	r.Use(static.Serve("/", static.LocalFile("dist", true)))
	r.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := ioutil.ReadFile("dist/index.html")
			if (err) != nil {
				c.Writer.WriteHeader(404)
				c.Writer.WriteString("Not Found")
				return
			}
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Accept", "text/html")
			c.Writer.Write((content))
			c.Writer.Flush()
		}
	})
	r.Use(Cors())
	user_route := r.Group("/api/user")
	{

		user_route.POST("/login",user.User_Api.Login_POST) 
		

	}


	admin.Admin_User_api.Admin_Create()
	live_route := r.Group("/api/live")
	live_route.Use(admin.Handler_Admin_auth())
	{
		live_route.GET("/up/list",live.Live_api.Live_List_get)
		live_route.POST("/up/list",live.Live_api.Live_List_post)
		live_route.PUT("/up/list",live.Live_api.Live_List_put)
		live_route.DELETE("/up/list",live.Live_api.Live_List_del)
		live_route.GET("/video/list",live.Live_api.Live_Video_list_get)
		live_route.POST("/video/list",live.Live_api.Live_Video_list_post)
		live_route.PUT("/video/list",live.Live_api.Live_Video_list_put)
		live_route.DELETE("/video/list",live.Live_api.Live_Video_list_del)
		live_route.GET("/video/down",live.Live_api.Live_Video_down)
	}
	conf_route := r.Group("/api/conf")
	conf_route.Use(admin.Handler_Admin_auth())
	{
		conf_route.GET("/set",liveset.Live_Set_get)
		conf_route.PUT("/set",liveset.Live_Set_put)
		conf_route.GET("/ping",liveset.Live_ping)
	}







	ip := conf.IniConf.Section("server").Key("ip").String() + ":" +conf.IniConf.Section("server").Key("port").String() 

    sysType := runtime.GOOS
    if sysType == "linux" {
    } 
    if sysType == "windows" {
		exec.Command(`cmd`, `/c`, `start`, fmt.Sprintf("http://%v/fc2/up",ip)).Start()
    }
	
	log.Println("程序启动完成")
	r.Run(ip) 
}