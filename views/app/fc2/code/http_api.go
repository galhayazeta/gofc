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
	Proxy_ip string      
	Timeout     int        
	ConnTimeout  int     
	MaxConnPerHost int      
}

func (obj *Client_Conf)ClientInit() *http.Client {
	if obj.Timeout == 0{
		obj.Timeout=6
	}
	if obj.ConnTimeout == 0{
		obj.ConnTimeout= 3
	}
	if obj.MaxConnPerHost == 0{
		obj.MaxConnPerHost = 200
	}

	var t *http.Transport
	if obj.Proxy_ip == "" {
		t = &http.Transport{
			
			MaxIdleConnsPerHost: obj.MaxConnPerHost,
			MaxConnsPerHost:     0,
			IdleConnTimeout:     time.Duration(obj.ConnTimeout) * time.Minute,
		}
	} else {
		
		u, _ := url.Parse(obj.Proxy_ip)

		t = &http.Transport{
			
			MaxIdleConnsPerHost: obj.MaxConnPerHost,
			MaxConnsPerHost:     0,
			IdleConnTimeout:     time.Duration(obj.ConnTimeout) * time.Minute,
			Proxy:               http.ProxyURL(u),
		}
	}
	return &http.Client{
		Transport: t,
		Timeout:   time.Duration(obj.Timeout) * time.Second,
	} 
}


func (obj *Http_Api_Class) req_send(req *http.Request) ([]byte, *http.Response, error) {
	var body []byte
	var resp *http.Response

	resp, err := obj.Client.Do(req) 
	if err != nil {
		fmt.Println("err:",err)
		return body, resp, errors.New("网络请求失败")
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close() 
		return body, resp, err 
		
	}
	resp.Body.Close() 
	return body, resp, nil
}


func (obj *Http_Api_Class) Get_Text(link_url string, head map[string]string) (string, *http.Response, error) {
	
	req, _ := http.NewRequest("GET", link_url, nil) 
	
	for key, val := range head { 
		req.Header.Set(key, val)
	}
	body, res, err := obj.req_send(req)
	if err != nil {
		return "", res, err
	}
	return string(body), res, nil 
}


func (obj *Http_Api_Class) Get_Byte(link_url string, head map[string]string) ([]byte, *http.Response, error) {
	
	var body []byte

	req, _ := http.NewRequest("GET", link_url, nil) 
	
	for key, val := range head { 
		req.Header.Set(key, val)
	}
	body, res, err := obj.req_send(req)

	if err != nil {
		return body, res, err
	}
	return body, res, nil 
}


func (obj *Http_Api_Class) Get_Json(link_url string, head map[string]string) ([]byte, *http.Response, error) {
	
	req, _ := http.NewRequest("GET", link_url, nil) 
	
	for key, val := range head { 
		req.Header.Set(key, val)
	}
	body, res, err := obj.req_send(req)
	if err != nil {
		return body, res, err
	}

	return body, res, nil

}


func (obj *Http_Api_Class) Post_From_Json(link_url string,  head map[string]string , data map[string]interface{}) ([]byte, *http.Response, error) {
	body := []byte{}
	urlValues := url.Values{}    
	for key, val := range data { 
		urlValues.Add(key, val.(string)) 
	}
	req, _ := http.NewRequest("POST", link_url, strings.NewReader(urlValues.Encode())) 
	
	for key, val := range head { 
		req.Header.Set(key, val)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded") 
	body, res, err := obj.req_send(req)
	if err != nil {
		return body, res, err
	}
	return body,res, nil
}


func (obj *Http_Api_Class) Post_Json_Json(link_url string , head map[string]string , data map[string]interface{}) (map[string]interface{}, *http.Response, error) {

	var r map[string]interface{}
	r = make(map[string]interface{})

	respdata, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("err TextGet1: %v\n", err)
		return r, nil, errors.New("提交参数转换为json失败")
	}

	req, _ := http.NewRequest("POST", link_url, bytes.NewReader(respdata)) 

	
	for key, val := range head { 
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

	return r, res, nil 

}



