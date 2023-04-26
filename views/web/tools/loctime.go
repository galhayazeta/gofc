package tools

import "time"

var Time_zone  *time.Location;

func init(){
	var err error;
	Time_zone, err = time.LoadLocation("Asia/Shanghai") 
	if err != nil {
		Time_zone = time.FixedZone("CST", 8*3600)
	}
}


