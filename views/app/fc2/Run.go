package fc2

import (
	"api/views/app/fc2/code"
	"api/views/logs"
	"api/views/web/database"
	"api/views/web/tools"
	"fmt"
	"log"
	"time"
)

func Fc2_live_Start(uid string) {

	p := database.Live_Conf_Datas_models{}
	database.DB.Where("website = ?", "fc2").First(&p)

	var api_conf = &code.Client_Conf{
		Proxy_ip:       p.Proxy,
		Timeout:        10,
		ConnTimeout:    3,
		MaxConnPerHost: 200,
	}
	api := &code.Http_Api_Class{
		Client: api_conf.ClientInit(),
	}
	info_api := code.Info_class{
		Uid: uid,
		Api: api,
	}
	title, _, name, wss_url, err := info_api.Info_run()
	if err != nil {
		logs.Log_write("录播_正常","Fc2_live_Start",fmt.Sprintf("获取wss链接失败:%v",err))
	}


	w := code.Ws_class{
		Link:       wss_url,
		Proxy:      p.Proxy,
		Quality:    p.Quality,
		Uid:        uid,
		Name:       name,
		VideoName:  title,
		State:      true,
		Start_time: time.Now().In(tools.Time_zone).Format("2006-01-02-15_04_05"),
	}
	w.Wss_conn()

	for i := 0; i < 3; i++ {
		logs.Log_write("录播_正常","Fc2_live_Start",fmt.Sprintf("避免网络原因造成的中断 开始进行重试:第%v次",i))
		_, _, _, wss_url, err := info_api.Info_run()
		if err != nil {
			logs.Log_write("录播_正常","Fc2_live_Start",fmt.Sprintf("获取wss链接失败:%v",err))
			continue
		}
		w.Link = wss_url
		w.State = true
		if w.Video_Down_state {
			w.Video_Down.State = true
		}

		w.Wss_conn()
	}
	if w.Video_Down_state {
		logs.Log_write("录播_正常","Fc2_live_Start",fmt.Sprintf("录播完成 开始合并视频"))
		video_path := w.Video_Down.Video_hb()
		w.Video_Down.Video_size_judge()
		w.Video_Down.Video_info_post(video_path)
	}


	user := database.Up_Datas_models{}
	database.DB.Where("uid = ?", uid).First(&user)
	user.State = "监控中"
	database.DB.Save(&user)
	api.Client.CloseIdleConnections()

}

func Fc2_main() {

	
	for {
		p := []database.Up_Datas_models{}
		
		database.DB.Where("website = ? AND  state = ?", "fc2", "监控中").Find(&p)

		open_dic, err := code.Fc2_open_live()
		if err != nil {
			logs.Log_write("录播_错误","Fc2_main",fmt.Sprintf("网站开播信息读取错误 :%v",err))
		}
	
		log.Println("检测中")
		for _, v := range p {
			_, ok := open_dic[v.Uid]
			if !ok {
				continue
			}
		
			go Fc2_live_Start(v.Uid)
		
			v.State = "录播中"
			database.DB.Save(&v)
		}
		
		time.Sleep(time.Second * 60)
	}

}
