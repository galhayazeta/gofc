package code

import (
	"api/views/logs"
	"api/views/web/database"
	"encoding/json"
	"fmt"
)



type Info_class struct{
	Proxy  string
	Uid string 
	Api *Http_Api_Class
} 


type channel_json_struct struct {
	Status int `json:"status"`
	Data   struct {
		ChannelData struct {
			Channelid string `json:"channelid"`
			Userid    string `json:"userid"`
			Adult     int    `json:"adult"`
			Twoshot   int    `json:"twoshot"`
			Title     string `json:"title"`
			Info      string `json:"info"`
			Image     string `json:"image"`
			LoginOnly int    `json:"login_only"`
			GiftLimit int    `json:"gift_limit"`
			GiftList  []struct {
				ID   int      `json:"id"`
				Type int      `json:"type"`
				URL  []string `json:"url"`
				Name string   `json:"name"`
			} `json:"gift_list"`
			CommentLimit       string `json:"comment_limit"`
			Tfollow            int    `json:"tfollow"`
			Tname              string `json:"tname"`
			Fee                int    `json:"fee"`
			Amount             int    `json:"amount"`
			Interval           int    `json:"interval"`
			Category           string `json:"category"`
			CategoryName       string `json:"category_name"`
			IsOfficial         int    `json:"is_official"`
			IsPremiumPublisher int    `json:"is_premium_publisher"`
			IsLinkShare        int    `json:"is_link_share"`
			Ticketid           int    `json:"ticketid"`
			IsPremium          int    `json:"is_premium"`
			TicketPrice        int    `json:"ticket_price"`
			TicketOnly         int    `json:"ticket_only"`
			IsApp              int    `json:"is_app"`
			IsVideo            int    `json:"is_video"`
			IsRest             int    `json:"is_rest"`
			Count              int    `json:"count"`
			IsPublish          int    `json:"is_publish"`
			IsLimited          int    `json:"is_limited"`
			Start              int64  `json:"start"`
			Version            string `json:"version"`
			Fc2Channel         struct {
				Result      int           `json:"result"`
				Userid      int           `json:"userid"`
				Fc2ID       int           `json:"fc2id"`
				Adult       int           `json:"adult"`
				Title       string        `json:"title"`
				Description string        `json:"description"`
				URL         string        `json:"url"`
				Images      []interface{} `json:"images"`
			} `json:"fc2_channel"`
			ControlTag          string `json:"control_tag"`
			PublishMethod       string `json:"publish_method"`
			VideoStereo3D       int    `json:"video_stereo3d"`
			VideoMapping        int    `json:"video_mapping"`
			VideoHorizontalView int    `json:"video_horizontal_view"`
		} `json:"channel_data"`
		ProfileData struct {
			Userid string `json:"userid"`
			Fc2ID  string `json:"fc2id"`
			Name   string `json:"name"`
			Info   string `json:"info"`
			Icon   string `json:"icon"`
			Image  string `json:"image"`
			Sex    string `json:"sex"`
			Age    string `json:"age"`
		} `json:"profile_data"`
		UserData struct {
			IsLogin       int    `json:"is_login"`
			Userid        int    `json:"userid"`
			Fc2ID         int    `json:"fc2id"`
			Icon          string `json:"icon"`
			Name          string `json:"name"`
			Point         int    `json:"point"`
			AdultAccess   int    `json:"adult_access"`
			Recauth       int    `json:"recauth"`
			IsPremiumUser int    `json:"is_premium_user"`
			GiftList      []struct {
				ID       int `json:"id"`
				Category int `json:"category"`
				Amount   int `json:"amount"`
			} `json:"gift_list"`
			Stamina struct {
				Timestamp int             `json:"timestamp"`
				Stamina   [][]interface{} `json:"stamina"`
			} `json:"stamina"`
		} `json:"user_data"`
	} `json:"data"`
}



func (obj *Info_class)Channel_Get()(channel_version,title,userid,name string,err error) {
	link := "https://live.fc2.com/api/memberApi.php"

	head := make(map[string]string, 0)
	head["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36"
	head["authority"] = "live.fc2.com.com"
	head["referer"] = fmt.Sprintf("https://live.fc2.com/%s/",obj.Uid)
	data := make(map[string]interface{}, 0)
	data["channel"] = "1"
	data["profile"] = "1"
	data["user"] = "1"
	data["streamid"] = obj.Uid


	json_data := channel_json_struct{}
	body, _, err := obj.Api.Post_From_Json(link, head, data)
	if err !=nil{
		logs.Log_write("录播_错误","Channel_Get",fmt.Sprintf("发送网络请求失败:%v",err))
		return channel_version,title,userid,name,err
	}
	err = json.Unmarshal(body, &json_data)
	if err != nil{
		logs.Log_write("录播_错误","Channel_Get",fmt.Sprintf("未开播 数据格式转换错误 str --> json  原因为:%v",err))
		return channel_version,title,userid,name,err
	}
	channel_version=json_data.Data.ChannelData.Version   // ["data"]["channel_data"]["version"]
	title=json_data.Data.ChannelData.Title  //["data"]["channel_data"]["title"]
	userid=json_data.Data.ChannelData.Userid  //["data"]["channel_data"]["userid"]
	name=json_data.Data.ProfileData.Name      //["data"]["profile_data"]["name"]
	return channel_version,title,userid,name,nil



}




type Control_token_json_struct struct {
	URL          string `json:"url"`
	Orz          string `json:"orz"`
	OrzRaw       string `json:"orz_raw"`
	ControlToken string `json:"control_token"`
	Status       int    `json:"status"`
}

func (obj *Info_class)Control_token_get(channel_version string) (string,error){
	link := "https://live.fc2.com/api/getControlServer.php"
	head := make(map[string]string, 0)
	head["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36"
	head["authority"] = "live.fc2.com.com"
	head["referer"] = fmt.Sprintf("https://live.fc2.com/%s/",obj.Uid)
	data := make(map[string]interface{}, 0)
	data["channel_id"]=obj.Uid
	data["mode"] = "play"
	data["orz"]=""
	data["channel_version"]=channel_version
	data["client_version"]="2.1.1  [1]"
	data["client_type"]="pc"
	data["client_app"]="browser_hls"
	data["ipv6"]=""
	data["comment"]="1"


	json_data := Control_token_json_struct{}
	body, _, err := obj.Api.Post_From_Json(link, head, data)
	if err !=nil{
		logs.Log_write("录播_错误","Control_token_get",fmt.Sprintf("发送网络请求失败:%v",err))
		return "",err
	}
	err = json.Unmarshal(body, &json_data)
	if err != nil{
		logs.Log_write("录播_错误","Control_token_get",fmt.Sprintf("未开播 数据格式转换错误 str --> json  原因为:%v",err))
		return "",err
	}
	wss_url := json_data.URL+"?control_token="+json_data.ControlToken
	return wss_url,nil

}

func (obj *Info_class)Info_run() (string,string,string,string,error){
	channel_version, title, userid, name, err := obj.Channel_Get()
	if err !=nil{
		logs.Log_write("录播_正常","Info_run",fmt.Sprintf("未开播:%v",err))
		return channel_version, title, userid, name,err
	}
	wss_url, err := obj.Control_token_get(channel_version)
	if err !=nil{
		logs.Log_write("录播_正常","Info_run",fmt.Sprintf("未开播:%v",err))
		return channel_version, title, userid, name,err
	}
	return title, userid, name,wss_url,err

}


type FC2_open_live_json struct {
	Link    string `json:"link"`
	IsAdult int    `json:"is_adult"`
	Time    int    `json:"time"`
	Channel []struct {
		ID             string `json:"id"`
		Bid            string `json:"bid"`
		Video          int    `json:"video"`
		App            int    `json:"app"`
		Category       int    `json:"category"`
		Type           int    `json:"type"`
		Fc2ID          int    `json:"fc2id"`
		Name           string `json:"name"`
		Title          string `json:"title"`
		Image          string `json:"image"`
		Start          string `json:"start"`
		StartTime      int64  `json:"start_time"`
		Sex            string `json:"sex"`
		Pay            int    `json:"pay"`
		Interval       int    `json:"interval"`
		Amount         int    `json:"amount"`
		Lang           string `json:"lang"`
		Total          int    `json:"total"`
		Count          int    `json:"count"`
		Login          int    `json:"login"`
		CommentL       int    `json:"comment_l"`
		Tid            int    `json:"tid"`
		Price          int    `json:"price"`
		Official       int    `json:"official"`
		CommentScore   int    `json:"comment_score"`
		DenyCountryFlg string `json:"deny_country_flg"`
		Panorama       int    `json:"panorama"`
	} `json:"channel"`
}


func Fc2_open_live()(result map[string]string, err error){
	//获取开播的uid
	result=map[string]string{}
	link := "https://live.fc2.com/adult/contents/allchannellist.php"
	head := make(map[string]string, 0)
	head["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36"
	head["authority"] = "live.fc2.com.com"
	head["referer"] = "https://live.fc2.com/adult/"
	head["origin"] = "https://live.fc2.com"

	p:=database.Live_Conf_Datas_models{}
	database.DB.Where("website = ?","fc2").First(&p)
	var api_conf = &Client_Conf{
		Proxy_ip:       p.Proxy,
		Timeout:        20,
		ConnTimeout:    3,
		MaxConnPerHost: 200,
	}

	api:=&Http_Api_Class{
		Client :api_conf.ClientInit(),
	}


	body, _, err := api.Get_Json(link, head)
	if err != nil{
		logs.Log_write("录播_错误","Fc2_open_live",fmt.Sprintf("网络请求失败:%v",err))
		return result,err
	}
	r:=FC2_open_live_json{}
	err = json.Unmarshal(body, &r)
	if err != nil{
		logs.Log_write("录播_错误","Fc2_open_live",fmt.Sprintf("开播信息转换为json格式错误:%v",err))
		return result,err
	}	

	for _,v := range r.Channel{
		if v.Login ==0{
			result[v.ID] = "1"
		}
	}
	return result,nil
}
