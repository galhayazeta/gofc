package liveset

import (
	"api/views/web/database"
	"api/views/web/tools"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)



func Live_Set_get(req *gin.Context){

	website := req.DefaultQuery("website","fc2")
	p := database.Live_Conf_Datas_models{}
	database.DB.Where("website = ?",website).First(&p)
	req.JSON(tools.Response(200, "查询成功", p))
	return
}

func Live_ping(req *gin.Context){

	proxy := req.DefaultQuery("proxy","")
	

	var t *http.Transport
	if proxy == "" {
		t = &http.Transport{
			
			MaxIdleConnsPerHost: 1,
			MaxConnsPerHost:     0,
			IdleConnTimeout:     time.Duration(3) * time.Minute,
		}
	} else {

		u, _ := url.Parse(proxy)
		t = &http.Transport{
			
			MaxIdleConnsPerHost: 1,
			MaxConnsPerHost:     0,
			IdleConnTimeout:     time.Duration(3) * time.Minute,
			Proxy:               http.ProxyURL(u),
		}
	}
	test_api:=&http.Client{
		Transport: t,
		Timeout:   time.Duration(20) * time.Second,
	}
	r, _ := http.NewRequest("GET", "https://live.fc2.com/adult/contents/allchannellist.php", nil) 
	r.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
	r2, err := test_api.Do(r)
	

	t.Clone()
	if err !=nil{
		req.JSON(tools.Response(503, "测试连通性失败", err))
		return
	}
	if r2.StatusCode == 200{
		req.JSON(tools.Response(200, "测试连通性成功", ""))
		return
	}else{
		req.JSON(tools.Response(503, "测试连通性失败", err))
		return
	}
}

type Live_conf_json struct{
	Proxy string         `json:"proxy"`   
	Quality string      `json:"quality"`   
	Website string        `json:"website"`   

}

func Live_Set_put(req *gin.Context){

	json_data := Live_conf_json{}
	err := req.ShouldBindJSON(&json_data)
	if err !=nil{
		req.JSON(tools.Response(503, "参数格式转换错误", err))
		return
	}
	p := database.Live_Conf_Datas_models{}
	database.DB.Where("website = ?",json_data.Website).First(&p)
	p.Proxy = json_data.Proxy
	p.Quality = json_data.Quality
	database.DB.Save(&p)
	req.JSON(tools.Response(200, "修改成功", ""))
	return
}