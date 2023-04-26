package live


type Up_info_add struct{
	Name string      `json:"name"`     
	Uid  string         `json:"uid"`      
	State string        `json:"state"`      
	Website string     `json:"website"`       
}


type Up_info_put struct{
	Id uint      `json:"id"`           
	Name string      `json:"name"`     
	Uid  string         `json:"uid"`      
	State string        `json:"state"`      
	Website string     `json:"website"`       
}



type Video_info_add struct{
	Video_name string    `json:"video_name"`      
	Uid string          `json:"uid" `       
	Name string          `json:"name"`    
	Size string          `json:"size" `      
	Website string        `json:"website"`    
	Path string          `json:"path"`    
}



type Video_info_put struct{
	Id  uint      `json:"id"`
	Video_name string    `json:"video_name"`      
	Uid string          `json:"uid" `       
	Name string          `json:"name"`    
	Size string          `json:"size"`      
	Website string        `json:"website"`    
	Path string          `json:"path"`    
}

