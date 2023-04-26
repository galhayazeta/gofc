package code

import (
	"api/views/logs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type Ws_class struct {
	Link        string   
	Proxy       string   
	Heart_count int     
	Heart_time  time.Time 
	Quality     string   
	Uid         string    
	Name        string   
	VideoName   string   
	Start_time  string    

	Conn             *websocket.Conn   
	Retry_count      int              
	State            bool             
	Video_Down       *Video_Down_class 
	Video_Down_state bool            
}


func (obj *Ws_class) Wss_conn() {
	var head = &http.Header{}

	head.Set("authority", "live.fc2.com.com")
	head.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36")

	var dl websocket.Dialer
	if obj.Proxy == "" {
		dl = websocket.Dialer{}
	} else {
		u, _ := url.Parse(obj.Proxy)
		dl = websocket.Dialer{
			Proxy: http.ProxyURL(u),
		} 
	}

	conn, _, err := dl.Dial(obj.Link, *head) 
	if err != nil {
	
		logs.Log_write("录播_错误", "Wss_conn", fmt.Sprintf("建立链接失败 :%v", err))
		return
	}
	obj.Conn = conn


	go obj.msg_receive()

	obj.heart_beat_send()

	
	err = obj.video_link_send()
	if err != nil {
		logs.Log_write("录播_错误", "Wss_conn", fmt.Sprintf("视频连接获取失败 :%v", err))
		return
	}


	for {
		if !obj.State {
			log.Println("连接结束")
			return
		}



		d := time.Now().Sub(obj.Heart_time).Seconds()
		if d > 30 {
			err := obj.heart_beat_send()
			if err != nil {
		
				obj.State = false
				logs.Log_write("录播_正常", "Wss_conn", fmt.Sprintf("心跳包发送失败 终止连接 :%v", err))
				obj.Conn.Close()
				if obj.Video_Down_state {
					obj.Video_Down.State = false
				}
				return
			}
		}
		time.Sleep(time.Millisecond * 500)

		if obj.Video_Down_state {
			if !obj.Video_Down.State{
				log.Println("通过obj.Video_Down.State 进行连接结束")
				return
			}
		}


	}
}


func (obj *Ws_class) msg_receive() {
	for {
		_, p, err := obj.Conn.ReadMessage()
		if err != nil {
			
			logs.Log_write("录播_正常", "msg_receive", fmt.Sprintf("获取消息失败 关闭连接:%v", err))
			obj.Conn.Close()
			if obj.Video_Down_state {
				obj.Video_Down.State = false
			}
		
			return
		}
		obj.msg_code_check(string(p))
	}

}


func (obj *Ws_class) heart_beat_send() error {
	msg := `{"name":"heartbeat","arguments":{},"id":` + strconv.Itoa(obj.Heart_count) + "}"
	err := obj.Conn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		logs.Log_write("录播_正常", "heart_beat_send", fmt.Sprintf("发送心跳请求包失败:%v", err))
		return err
	}
	obj.Heart_time = time.Now()

	obj.Heart_count++
	return nil

}


func (obj *Ws_class) video_link_send() error {
	err := obj.Conn.WriteMessage(websocket.TextMessage, []byte(`{"name":"get_hls_information","arguments":{},"id":1}`))
	if err != nil {
		logs.Log_write("录播_正常", "video_link_send", fmt.Sprintf("发送视频链接请求包失败:%v", err))
		return err
	}
	return nil

}


func (obj *Ws_class) msg_code_check(data string) {

	if strings.Contains(data, "playlists") {
		m3u8link, err := obj.quality_select(data)
		if err != nil {
			logs.Log_write("录播_错误", "msg_code_check", fmt.Sprintf("获取视频链接出现错误:%v", err))
			return
		}
	
		video_down := &Video_Down_class{
			M3u8link:   m3u8link,
			Proxy:      obj.Proxy,
			Uid:        obj.Uid,
			Name:       obj.Name,
			VideoName:  obj.VideoName,
			Start_time: obj.Start_time,
			State:      true,
			Index_map:  map[string]string{},
			Count: 0,
		}
		obj.Video_Down = video_down
		obj.Video_Down_state = true
		go obj.Video_Down.Video_run()
		return
	}

	if !strings.Contains(data, "code") {
		return
	}
	r := regexp.MustCompile(`"code":(\d+)`)
	result := r.FindAllStringSubmatch(data, -1)
	if len(result) == 0 {
		return
	}
	s := result[0][1]
	fmt.Println("code:", s)
	switch s {
	case "4101":
		logs.Log_write("录播_正常", "msg_code_check", fmt.Sprintf("主播开启收费"))
	case "4507":
		logs.Log_write("录播_正常", "msg_code_check", fmt.Sprintf("主播开启收费"))
	case "4000":
		logs.Log_write("录播_正常", "msg_code_check", fmt.Sprintf("节目播放完毕"))
	case "4001":
		logs.Log_write("录播_正常", "msg_code_check", fmt.Sprintf("节目播放完毕"))
	case "4512":
		logs.Log_write("录播_正常", "msg_code_check", fmt.Sprintf("开启了另外一个连接 网站不允许多连接"))
	case "4114":
		logs.Log_write("录播_正常", "msg_code_check", fmt.Sprintf("4114暂时不知道是什么报错"))
	default:
		logs.Log_write("录播_正常", "msg_code_check", fmt.Sprintf("未知状态码:%v",s))
	}

	obj.State = false
	obj.Conn.Close()
	if obj.Video_Down_state {
		obj.Video_Down.State = false
	}
	return
}

type Quality_json struct {

	Name      string `json:"name"`
	ID        int    `json:"id"`
	Arguments struct {
		Status    int `json:"status"`
		Playlists []struct {
			Mode   int    `json:"mode"`
			Status int    `json:"status"`
			URL    string `json:"url"`
		} `json:"playlists"`
		PlaylistsHighLatency []struct {
			Mode   int    `json:"mode"`
			Status int    `json:"status"`
			URL    string `json:"url"`
		} `json:"playlists_high_latency"`
		PlaylistsMiddleLatency []struct {
			Mode   int    `json:"mode"`
			Status int    `json:"status"`
			URL    string `json:"url"`
		} `json:"playlists_middle_latency"`
	} `json:"arguments"`
}


func (obj *Ws_class) quality_select(data string) (string, error) {

	video_quality := make(map[string]int)
	video_quality["150Kbps"] = 11
	video_quality["400Kbps"] = 21
	video_quality["1.2Mbps"] = 31
	video_quality["2Mbps"] = 41
	video_quality["3Mbps"] = 51
	video_quality["sound"] = 91

	json_data := Quality_json{}
	if err := json.Unmarshal([]byte(data), &json_data); err != nil {
		logs.Log_write("录播_严重", "quality_select", fmt.Sprintf("json格式转换错误:%v",err))
		return "", err
	}
	s := json_data.Arguments.PlaylistsHighLatency
	var m3u8link string
	for _, v := range s {
		if v.Mode == video_quality[obj.Quality] {
			logs.Log_write("录播_正常", "quality_select", fmt.Sprintf("初始画质选择成功:%v", obj.Quality))
			m3u8link = v.URL
		}
	}
	default_list := []string{"3Mbps", "2Mbps", "1.2Mbps", "400Kbps", "150Kbps"}
	st := false
	if m3u8link == "" {
		
		for _, v := range default_list {
			if st {
				break
			}
			i := video_quality[v]
			for _, qual := range s {
				if qual.Mode == i {
					logs.Log_write("录播_正常", "quality_select", fmt.Sprintf("默认画质选择成功:%v", v))
					m3u8link = qual.URL
					st = true
					break
				}
			}
		}
	}
	return m3u8link, nil

}
