package code

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Http_Api_Class struct {
	Client *http.Client
}

type Client_Conf struct {
	Proxy_ip string        //代理地址
	Timeout     int         //响应超时时间    单位/秒
	ConnTimeout  int       //空闲连接在连接池中保留的时长   单位/分钟
	MaxConnPerHost int        //最大空闲连接数
}

func (obj *Client_Conf)ClientInit() *http.Client {
	//初始化

	//设置默认配置
	if obj.Timeout == 0{
		obj.Timeout=6
	}
	if obj.ConnTimeout == 0{
		obj.ConnTimeout= 3
	}
	if obj.MaxConnPerHost == 0{
		obj.MaxConnPerHost = 200
	}


	//配置连接参数
	var t *http.Transport
	if obj.Proxy_ip == "" {
		t = &http.Transport{
			// MaxIdleConns:    0,
			MaxIdleConnsPerHost: obj.MaxConnPerHost,
			MaxConnsPerHost:     0,
			IdleConnTimeout:     time.Duration(obj.ConnTimeout) * time.Minute,
		}
	} else {
		//进行代理转换
		u, _ := url.Parse(obj.Proxy_ip)

		t = &http.Transport{
			// MaxIdleConns:    0,
			MaxIdleConnsPerHost: obj.MaxConnPerHost,
			MaxConnsPerHost:     0,
			IdleConnTimeout:     time.Duration(obj.ConnTimeout) * time.Minute,
			Proxy:               http.ProxyURL(u),
		}
	}
	return &http.Client{
		Transport: t,
		Timeout:   time.Duration(obj.Timeout) * time.Second,
	} //创建一个连接
}

// 请求函数复用 封装
func (obj *Http_Api_Class) req_send(req *http.Request) ([]byte, *http.Response, error) {
	var body []byte
	var resp *http.Response

	resp, err := obj.Client.Do(req) //设置完毕后 发出网络请求
	if err != nil {
		fmt.Println("err:",err)
		return body, resp, errors.New("网络请求失败")
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close() //关闭本次连接
		return body, resp, err // errors.New("数据转换为字符串错误 一般是返回的数据是压缩的 需要在请求头里面设置接口明码返回")
		//	设置接口明码返回  head["Accept-Encoding"] = "identity"
	}
	resp.Body.Close() //关闭本次连接
	return body, resp, nil
}

// 普通get请求 返回字符串  用于爬取网页
func (obj *Http_Api_Class) Get_Text(link_url string, head map[string]string) (string, *http.Response, error) {
	//普通get请求
	req, _ := http.NewRequest("GET", link_url, nil) //配置本次请求的方式和url
	//设置请求头内容
	for key, val := range head { //遍历map
		req.Header.Set(key, val)
	}
	body, res, err := obj.req_send(req)
	if err != nil {
		return "", res, err
	}
	return string(body), res, nil //将数据转换为字符串类型 并且作为函数的返回值
}

// 普通get请求 返回数据流 用于下载文件
func (obj *Http_Api_Class) Get_Byte(link_url string, head map[string]string) ([]byte, *http.Response, error) {
	//普通get请求 返回数据流 用于下载文件
	var body []byte

	req, _ := http.NewRequest("GET", link_url, nil) //配置本次请求的方式和url
	//设置请求头内容
	for key, val := range head { //遍历map
		req.Header.Set(key, val)
	}
	body, res, err := obj.req_send(req)

	if err != nil {
		return body, res, err
	}
	return body, res, nil // 并且作为函数的返回值
}

// 普通get请求 返回json数据
func (obj *Http_Api_Class) Get_Json(link_url string, head map[string]string) ([]byte, *http.Response, error) {
	//普通get请求 返回json数据
	req, _ := http.NewRequest("GET", link_url, nil) //配置本次请求的方式和url
	//设置请求头内容
	for key, val := range head { //遍历map
		req.Header.Set(key, val)
	}
	body, res, err := obj.req_send(req)
	if err != nil {
		return body, res, err
	}

	return body, res, nil

}

// // 普通的post 请求 返回网页
func (obj *Http_Api_Class) Post_From_Text(link_url string ,  head map[string]string , data map[string]interface{}) (string, *http.Response, error) {

	urlValues := url.Values{}    //创建提交参数表单
	for key, val := range data { //遍历map
		urlValues.Add(key, val.(string)) //将内容写入到表单内容
	}

	req, _ := http.NewRequest("POST", link_url, strings.NewReader(urlValues.Encode())) //配置本次请求的方式和url
	//设置请求头内容
	for key, val := range head { //遍历map
		req.Header.Set(key, val)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded") //通过设置请求头参数 让数据变成表单form类型进行提交  默认的提交的是data

	body, res, err := obj.req_send(req)
	if err != nil {
		return "", res, err
	}
	return string(body), res, nil //将数据转换为字符串类型 并且作为函数的返回值

}

// 普通的post请求 返回json
func (obj *Http_Api_Class) Post_From_Json(link_url string,  head map[string]string , data map[string]interface{}) ([]byte, *http.Response, error) {
	body := []byte{}
	urlValues := url.Values{}    //创建提交参数表单
	for key, val := range data { //遍历map
		urlValues.Add(key, val.(string)) //将内容写入到表单内容
	}
	req, _ := http.NewRequest("POST", link_url, strings.NewReader(urlValues.Encode())) //配置本次请求的方式和url
	//设置请求头内容
	for key, val := range head { //遍历map
		req.Header.Set(key, val)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded") //通过设置请求头参数 让数据变成表单form类型进行提交  默认的提交的是data
	body, res, err := obj.req_send(req)
	if err != nil {
		return body, res, err
	}
	return body,res, nil
}



// json的post请求  返回网页
func (obj *Http_Api_Class) Post_Json_Text(link_url string,  head map[string]string , data map[string]interface{}) (string, *http.Response, error) {

	respdata, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("err TextGet1: %v\n", err)
		return "", nil, errors.New("提交参数转换为json失败")
	}

	req, _ := http.NewRequest("POST", link_url, bytes.NewReader(respdata)) //配置本次请求的方式和url

	//设置请求头内容
	for key, val := range head { //遍历map
		req.Header.Set(key, val)
	}
	body, res, err := obj.req_send(req)
	if err != nil {
		return "", nil, err
	}
	return string(body), res, nil //将数据转换为字符串类型 并且作为函数的返回值
}

// json的post请求  返回json
func (obj *Http_Api_Class) Post_Json_Json(link_url string , head map[string]string , data map[string]interface{}) (map[string]interface{}, *http.Response, error) {

	var r map[string]interface{}
	r = make(map[string]interface{})

	respdata, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("err TextGet1: %v\n", err)
		return r, nil, errors.New("提交参数转换为json失败")
	}

	req, _ := http.NewRequest("POST", link_url, bytes.NewReader(respdata)) //配置本次请求的方式和url

	//设置请求头内容
	for key, val := range head { //遍历map
		req.Header.Set(key, val)
	}

	body, res, err := obj.req_send(req)
	if err != nil {
		return r, nil, err
	}

	err = json.Unmarshal([]byte(body), &r)
	if err != nil {
		fmt.Printf("err TextGet2: %v body %v\n", err)
		return r, nil, errors.New("数据转换为json格式失败")
	}

	return r, res, nil //将数据转换为字符串类型 并且作为函数的返回值

}








// func main() {
// 	// demo_http_get()
// 	// demo_http_post()

// 	// 使用示例
// 	api_conf := &Client_Conf{
// 		Proxy_ip:"",          //http://127.0.0.1:7890
// 		Timeout    :6,
// 		ConnTimeout: 3,
// 		MaxConnPerHost: 200,
// 	}
// 	api := &Http_Api_Class{
// 		Client: api_conf.ClientInit(), //http://127.0.0.1:7890
// 	}
// 	head:=make(map[string]string, 0)

// 	head["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36"

// 	// data := make(map[string]interface{}, 0)

// 	m, r, err := api.Get_Text("http://httpbin.org/get", head)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 		return
// 	}
// 	fmt.Printf("r.StatusCode: %v\n", r.StatusCode)
// 	fmt.Printf("m: %v\n", m)

// }
