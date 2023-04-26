package live

import (
	"api/views/web/database"
	"api/views/web/tools"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)


type Live_Api_struct struct {
	
}

var Live_api = Live_Api_struct{}




func (obj *Live_Api_struct)Live_List_get(req *gin.Context) {
	
	pagesize := req.DefaultQuery("pagesize", "")
	page := req.DefaultQuery("page", "")
	website := req.DefaultQuery("website", "")
	uid := req.DefaultQuery("uid", "")
	name := req.DefaultQuery("name", "")

	if pagesize == "" || page == "" {
		page = "1"
		pagesize = "10"
	}
	pagesize_int, err := strconv.Atoi(pagesize)
	if err != nil {
		req.JSON(tools.Response(503, "查询失败", "提交参数错误"))
		return
	}
	page_int, err := strconv.Atoi(page)
	if err != nil {
		req.JSON(tools.Response(503, "查询失败", "提交参数错误"))
		return
	}

	page_int = page_int - 1
	start := page_int * pagesize_int

	uplist := []database.Up_Datas_models{}

	var total int64

	if uid == "" && name == ""{
		database.DB.Model(&database.Up_Datas_models{}).Where("website = ?", website).Count(&total)
		database.DB.Where("website = ?", website).Order("updated_at desc").Limit(pagesize_int).Offset(start).Find(&uplist)
	
	}else{
		database.DB.Model(&database.Up_Datas_models{}).Where("website = ? AND (uid = ? OR name =? )", website, uid, name).Count(&total)
		database.DB.Where("website = ? AND (uid = ? OR name =? )", website, uid, name).Order("updated_at desc").Limit(pagesize_int).Offset(start).Find(&uplist)
	
	}

	data := gin.H{
		"total": total,
		"data":  uplist,
	}

	req.JSON(tools.Response(200, "查询成功", data))
	return
}

func (obj *Live_Api_struct)Live_List_post(req *gin.Context) {
	
	json_data := Up_info_add{}
	err := req.ShouldBindJSON(&json_data)
	if err != nil {
		req.JSON(tools.Response(503, "添加失败", "提交参数错误"))
		return
	}

	p := &database.Up_Datas_models{}



	
	d := database.DB.Where("website = ? AND (uid = ? OR name = ?)", json_data.Website, json_data.Uid, json_data.Name).First(&p)
	
	if d.RowsAffected != 0 {
		req.JSON(tools.Response(503, "添加失败", "up主姓名或uid已存在"))
		return
	}
	
	p = &database.Up_Datas_models{
		Name:    json_data.Name,
		Uid:     json_data.Uid,
		State:   json_data.State,
		Website: json_data.Website,
	}

	d = database.DB.Create(&p)
	if d.RowsAffected == 0 {
		req.JSON(tools.Response(503, "添加失败", fmt.Sprintf("原因为:%v", d.Error)))
		return
	}
	req.JSON(tools.Response(200, "添加成功", fmt.Sprintf("添加up:%v  uid:%v 成功", json_data.Name, json_data.Uid)))
	return

}

func (obj *Live_Api_struct)Live_List_put(req *gin.Context) {
	
	json_data := Up_info_put{}
	err := req.ShouldBindJSON(&json_data)
	if err != nil {
		req.JSON(tools.Response(503, "修改失败", "提交参数错误"))
		return
	}
	p := database.Up_Datas_models{}
	
	d := database.DB.First(&p, json_data.Id)
	if d.RowsAffected == 0 {
		req.JSON(tools.Response(503, "修改失败", "未查询到本条数据"))
		return
	}

	
	p.Name = json_data.Name
	p.Uid = json_data.Uid
	p.State = json_data.State
	p.Website = json_data.Website

	d = database.DB.Save(&p)
	if d.RowsAffected == 0 {
		req.JSON(tools.Response(503, "修改失败", fmt.Sprintf("更新数据时发生错误:%v", d.Error)))
		return
	}
	req.JSON(tools.Response(200, "修改成功", "修改成功"))
	return
}

func (obj *Live_Api_struct)Live_List_del(req *gin.Context) {
	
	
	id := req.DefaultQuery("id", "")
	if id == "" {
		req.JSON(tools.Response(503, "删除失败", "提交参数错误"))
		return
	}

	
	d := database.DB.Delete(&database.Up_Datas_models{}, id)
	if d.RowsAffected == 0 {
		req.JSON(tools.Response(503, "删除失败", fmt.Sprintf("删除数据时发生错误:%v", d.Error)))
		return
	}

	req.JSON(tools.Response(200, "删除成功", "删除成功"))
	return

}

func (obj *Live_Api_struct)Live_Video_list_get(req *gin.Context) {
	
	pagesize := req.DefaultQuery("pagesize", "")
	page := req.DefaultQuery("page", "")
	website := req.DefaultQuery("website", "")
	uid := req.DefaultQuery("uid", "")
	name := req.DefaultQuery("name", "")

	if pagesize == "" || page == "" {
		page = "1"
		pagesize = "10"
	}
	pagesize_int, err := strconv.Atoi(pagesize)
	if err != nil {
		req.JSON(tools.Response(503, "查询失败", "提交参数错误"))
		return
	}
	page_int, err := strconv.Atoi(page)
	if err != nil {
		req.JSON(tools.Response(503, "查询失败", "提交参数错误"))
		return
	}

	page_int = page_int - 1
	start := page_int * pagesize_int

	videolist := []database.Video_Datas_models{}

	var total int64
	
	if uid == "" && name == ""{
		database.DB.Model(&database.Video_Datas_models{}).Where("website = ?", website, uid, name).Count(&total)
		database.DB.Where("website = ?", website, uid, name).Limit(pagesize_int).Offset(start).Find(&videolist)
	
	}else{
		database.DB.Model(&database.Video_Datas_models{}).Where("website = ? AND (uid = ? OR name =? )", website, uid, name).Count(&total)
		database.DB.Where("website = ? AND (uid = ? OR name =? )", website, uid, name).Limit(pagesize_int).Offset(start).Find(&videolist)

	}

	data := gin.H{
		"total": total,
		"data":  videolist,
	}

	req.JSON(tools.Response(200, "查询成功", data))
	return

}

func (obj *Live_Api_struct)Live_Video_list_post(req *gin.Context) {
	
	json_data := Video_info_add{}
	err := req.ShouldBindJSON(&json_data)
	if err != nil {
		req.JSON(tools.Response(503, "添加失败", "提交参数错误"))
		return
	}

	p := database.Video_Datas_models{
		Video_name: json_data.Name,
		Uid:        json_data.Uid,
		Name:       json_data.Name,
		Size:       json_data.Size,
		Website:    json_data.Website,
		Path:       json_data.Path,
	}

	d := database.DB.Create(&p)
	if d.RowsAffected == 0 {
		req.JSON(tools.Response(503, "添加失败", fmt.Sprintf("原因为:%v", d.Error)))
		return
	}
	req.JSON(tools.Response(200, "添加成功", "视频添加成功"))
	return

}

func (obj *Live_Api_struct)Live_Video_list_put(req *gin.Context) {
	

	json_data := Video_info_put{}
	err := req.ShouldBindJSON(&json_data)
	if err != nil {
		req.JSON(tools.Response(503, "修改失败", "提交参数错误"))
		return
	}
	
	p := database.Video_Datas_models{}
	
	d := database.DB.First(&p, json_data.Id)
	if d.RowsAffected == 0 {
		req.JSON(tools.Response(503, "修改失败", "未查询到本条数据"))
		return
	}

	
	p.Video_name = json_data.Video_name
	p.Name = json_data.Name
	p.Uid = json_data.Uid
	p.Size = json_data.Size
	p.Path = json_data.Path
	p.Website = json_data.Website

	d = database.DB.Save(&p)
	if d.RowsAffected == 0 {
		req.JSON(tools.Response(503, "修改失败", fmt.Sprintf("更新数据时发生错误:%v", d.Error)))
		return
	}
	req.JSON(tools.Response(200, "修改成功", "修改成功"))
	return
}

func (obj *Live_Api_struct)Live_Video_list_del(req *gin.Context) {
	

	id := req.DefaultQuery("id", "")
	if id == "" {
		req.JSON(tools.Response(503, "删除失败", "提交参数错误"))
		return
	}
	
	p := database.Video_Datas_models{}
	d := database.DB.First(&p, id)
	if d.RowsAffected == 0 {
		req.JSON(tools.Response(503, "删除失败", fmt.Sprintf("删除数据时发生错误:%v", d.Error)))
		return
	}

	
	os.Remove(p.Path)


	
	d = database.DB.Delete(&database.Video_Datas_models{}, id)
	if d.RowsAffected == 0 {
		req.JSON(tools.Response(503, "删除失败", fmt.Sprintf("删除数据时发生错误:%v", d.Error)))
		return
	}

	req.JSON(tools.Response(200, "删除成功", "删除成功"))
	return


}

func (obj *Live_Api_struct)Live_Video_down(req *gin.Context) {
	
	id := req.DefaultQuery("id","")
	if id == "" {
		req.JSON(tools.Response(503, "提交参数错误", "提交参数错误"))
		return
	}
	p:= database.Video_Datas_models{}
	d := database.DB.First(&p, id)
	if d.RowsAffected == 0 {
		req.JSON(tools.Response(503, "未找到数据", "未找到数据"))
		return
	}


	filelen := len(strings.Split(p.Path, "/"))
	video_Name := ""
	if filelen <= 1 {
		filelen = len(strings.Split(p.Path, "\\"))
		video_Name = strings.Split(p.Path, "\\")[filelen-1]
	} else {
		video_Name = strings.Split(p.Path, "/")[filelen-1]
	}
	req.Header("Content-Type", "application/octet-stream")                
	req.Header("Content-Disposition", "attachment; filename="+video_Name) 
	req.Header("Content-Transfer-Encoding", "binary")                    
	req.File(p.Path)
	return

}

