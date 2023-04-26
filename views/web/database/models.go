package database

import "time"

func init() {
	Time_zone, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		Time_zone = time.FixedZone("CST", 8*3600)
	}
	time.Local = Time_zone 
}

type User_Datas_models struct {
	
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Username string `json:"username" gorm:"unique;"`  
	Password string `json:"password"`                 

	Level    int    `json:"level" gorm:"default:0;"`  
	State    string `json:"state" gorm:"default:正常;"` 
	UserImg  string `json:"userimg" gorm:""` 
}



type Up_Datas_models struct{
	
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name string        `json:"name" gorm:""`      
	Uid  string        `json:"uid" gorm:""`     
	State string        `json:"state" gorm:""`    
	Website string      `json:"website" gorm:""`    
	
}

type Video_Datas_models struct{
	
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Video_name string    `json:"video_name" gorm:""`      
	Uid string          `json:"uid" gorm:""`       
	Name string          `json:"name" gorm:""`    
	Size string          `json:"size" gorm:""`      
	Website string        `json:"website" gorm:""`    
	Path string          `json:"path" gorm:""`    

}


type Live_Conf_Datas_models struct{
	
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Proxy string         `json:"proxy" gorm:""`   
	Quality string      `json:"quality" gorm:""`   
	Website string        `json:"website" gorm:""`    

}
