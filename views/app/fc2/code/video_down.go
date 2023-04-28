package code

import (
	"api/views/logs"
	"api/views/web/database"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Video_Down_class struct {
	M3u8link   string            //m3u8连接
	Proxy      string            //代理
	Uid        string            //房间号
	Name       string            //up名
	VideoName  string            //视频名称
	Start_time string            //开始时间 字符串格式
	State      bool              //判断是否继续下载
	Http_api   *Http_Api_Class   //下载结构体
	Index_map   sync.Map        //map[string]string //储存已经下载的序列号



	Count  int        //避免网络波动过大 导致无限错误  超过一定数量 关闭下载
}

func (obj *Video_Down_class) m3u8_down() (result [][]string, err error) {
	//下载m3u8文件 并解析
	head := make(map[string]string)
	head["authority"] = "hls.live.fc2.com"
	head["accept"] = "*/*"
	head["accept-language"] = "zh-CN,zh;q=0.9,ja;q=0.8"
	head["origin"] = "https]=//live.fc2.com"
	head["referer"] = "https]=//live.fc2.com/"
	head["sec-fetch-dest"] = "empty"
	head["sec-fetch-mode"] = "cors"
	head["sec-fetch-site"] = "same-site"
	head["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36"

	s, _, err := obj.Http_api.Get_Text(obj.M3u8link, head)


	if err != nil {
		logs.Log_write("录播_错误","m3u8_down",fmt.Sprintf("m3u8_down 请求失败:%v",err))
		obj.Count++
		return
	}
	if s == "Forbidden." {
		logs.Log_write("录播_错误","m3u8_down",fmt.Sprintf("返回数据为Forbidden"))
		obj.State = false    //直播完毕 关闭录播
		return
	}
	r := regexp.MustCompile(`https:.*`)
	result = r.FindAllStringSubmatch(s, -1)


	return
}

func (obj *Video_Down_class) ts_down(ts_list [][]string) {
	//解析ts文件 并启动下载
	for _, v := range ts_list {
		if len(v) == 0 {
			return
		}
		ts_link := v[0]
		//使用正则获取 视频编号
		r := regexp.MustCompile(`(\d+\.ts)`)
		ts_index := r.FindAllStringSubmatch(ts_link, -1)
		if len(ts_index) == 0 {
			logs.Log_write("录播_错误","ts_down",fmt.Sprintf("未获取到 ts文件编号"))
			obj.Count++
			continue
		}
		//检测字典中是否已经下载过了
		_, ok := obj.Index_map.Load(ts_index[0][0]) //  [ts_index[0][0]]

		if ok {
			//下载过了 就不下载了
			continue
		}
		//未下载过 就开启下载
		go obj.ts_v_down(ts_link, ts_index[0][0])
	}
}

func (obj *Video_Down_class) ts_v_down(ts_link string, index string) {
	//下载视频到本地
	// fmt.Println("开始下载视频:", index)
	head := make(map[string]string)
	head["authority"] = "hls.live.fc2.com"
	head["accept"] = "*/*"
	head["accept-language"] = "zh-CN,zh;q=0.9,ja;q=0.8"
	head["origin"] = "https]=//live.fc2.com"
	head["referer"] = "https]=//live.fc2.com/"
	head["sec-fetch-dest"] = "empty"
	head["sec-fetch-mode"] = "cors"
	head["sec-fetch-site"] = "same-site"
	head["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36"



	b, r, err := obj.Http_api.Get_Byte(ts_link, head)


	if err != nil {
		logs.Log_write("录播_错误","ts_v_down",fmt.Sprintf("请求失败了:%v",err))
		obj.Count++
		return
	}
	
	if r.StatusCode != 200 {
		logs.Log_write("录播_错误","ts_v_down",fmt.Sprintf("状态码错误 状态码为:%v",r.StatusCode))
		obj.Count++
		return
	}

	//添加到已下载字典中
	obj.Index_map.Store(index,"已下载") // [index] = "已下载"

	//保存到本地
	obj.ts_save(b, index)

}

func (obj *Video_Down_class) ts_save(data []byte, fliename string) {
	//保存ts文件
	fileName := "./video/fc2/" + obj.Uid + "/" + obj.Start_time + "/" + fliename
	err := ioutil.WriteFile(fileName, data, 0644) // 保存文件
	if err != nil {
		logs.Log_write("录播_错误","ts_save",fmt.Sprintf("保存文件失败:%v",err))
		return
	}
}

func (obj *Video_Down_class) tital_mkdir() {
	//创建文件夹
	dirName := "./video/fc2/" + obj.Uid + "/" + obj.Start_time
	if _, err := os.Stat(dirName); os.IsNotExist(err) { // 检查文件夹是否存在
		err := os.MkdirAll(dirName, 0755) // 创建文件夹
		if err != nil {
			logs.Log_write("录播_严重","tital_mkdir",fmt.Sprintf("创建文件夹失败:%v",err))
			return
		}
		fmt.Println("文件夹创建成功")
	} else {
		fmt.Println("文件夹已存在")
	}
}

func (obj *Video_Down_class) Video_hb() string {
	//视频合并
	path := "./video/fc2/" + obj.Uid + "/" + obj.Start_time
	de, _ := os.ReadDir(path + "/") //获取所有需要合并的文件夹

	name_list := make([]int, 0, 5000) //创建一个用于储存文件名的切片
	for _, v := range de {            //将名称转换为数字类型
		s := strings.Replace(v.Name(), ".ts", "", -1)
		i, _ := strconv.Atoi(s)
		name_list = append(name_list, i)
	}
	sort.Ints(name_list) //进行排序

	//开始合并视频
	//创建一个新的文件
	path_top := filepath.Join(path, "../")
	f, err := os.OpenFile(path_top+"/"+obj.Uid+"_"+obj.Start_time+".ts", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)

	if err != nil {
		logs.Log_write("录播_严重","Video_hb",fmt.Sprintf("文件创建错误:%v",err))
		return ""
	}

	defer f.Close() //最后关闭文件
	w := bufio.NewWriter(f)
	for _, v := range name_list {
		path_name := path + "/" + fmt.Sprintf("%v", v) + ".ts"
		//开始合并视频
		buf1, err2 := os.ReadFile(path_name)
		if err2 != nil {
			logs.Log_write("录播_严重","Video_hb",fmt.Sprintf("文件块打开错误:%v",err))
			return ""
		}
		w.Write(buf1)
		w.Flush() //刷新缓冲区
	}

	err = os.RemoveAll(path) //删除无用文件夹
	if err != nil {
		logs.Log_write("录播_严重","Video_hb",fmt.Sprintf("删除文件夹失败:%v",err))
	}

	return path_top + "/" + obj.Uid + "_" + obj.Start_time + ".ts"
}

func (obj *Video_Down_class) Video_size_judge()int64 {
	//获取视频大小
	flie_path := "./video/fc2/"+ obj.Uid+"/" +obj.Uid + "_" + obj.Start_time + ".ts"

	fmt.Println("视频大小路径:",flie_path)
	fi, err := os.Stat(flie_path)
	if err != nil {
		log.Println("文件不存在")
		return 0
	}
	i := fi.Size() / 1024000 //文件的大小
	if i < 10 {
		//删除小于10M的视频
		err := os.Remove(flie_path)
		if err != nil {
			fmt.Printf("err%v\n", err)
		}
		logs.Log_write("录播_正常","Video_size_judge",fmt.Sprintf("删除小于10M的视频"))
	}
	logs.Log_write("录播_正常","Video_size_judge",fmt.Sprintf("视频大小为：%vMB",i))
	return i

}

func (obj *Video_Down_class) Video_info_post(video_path string) {
	//向数据库提交本次视频的信息
	log.Println("向数据库提交数据")
	i := obj.Video_size_judge()
	if i > 10{
		p :=database.Video_Datas_models{
			Video_name: obj.VideoName,
			Uid: obj.Uid,
			Name: obj.Name,
			Size: strconv.Itoa(int(i)) ,
			Website:"fc2",
			Path   :video_path,
		}
		database.DB.Create(&p)
	}

}

// 启动函数
func (obj *Video_Down_class) Video_run() {
	log.Println("启动视频下载函数")

	api_conf := &Client_Conf{
		Proxy_ip:       obj.Proxy,
		Timeout:        10,
		ConnTimeout:    3,
		MaxConnPerHost: 200,
	}
	obj.Http_api = &Http_Api_Class{
		Client: api_conf.ClientInit(),
	}

	obj.tital_mkdir()

	for {
		if !obj.State {
			logs.Log_write("录播_正常","Video_run",fmt.Sprintf("根据obj.State 停止下载"))
			break
		}
		if obj.Count > 5000{
			logs.Log_write("录播_严重","Video_run",fmt.Sprintf("网络波动过大 停止下载"))
			break
		}


		result, err := obj.m3u8_down()
		if err != nil {
			logs.Log_write("录播_正常","Video_run",fmt.Sprintf("获取m3u8错误 重试下载 :%v",err))
			continue
		}
		obj.ts_down(result)
		if !obj.State {
			logs.Log_write("录播_正常","Video_run",fmt.Sprintf("根据obj.State 停止下载"))
			break
		}
		time.Sleep(time.Second * 3)
	}

	obj.Http_api.Client.CloseIdleConnections()


}
