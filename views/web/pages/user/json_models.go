package user

import "time"



type Register_post_json struct{
	
	Username string `json:"username"`  
	Password string `json:"password"`    
	
	
}


type Login_post_json struct{
	
	Username string `json:"username"`  
	Password string `json:"password"`    
}


type User_info_result struct{
	
	CreatedAt time.Time `json:"createdat"`
	Username string `json:"username"`  
	Level    int    `json:"level"`  
	State    string `json:"state"` 
	UserImg  string `json:"userimg"` 
}

type User_put_json struct{
	
	UserImg  string `json:"userimg"` 


}