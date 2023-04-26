package tools

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct { 
	Username           string `json:"username"`
	Level              int    `json:"level"`
	jwt.StandardClaims       
}

var key = []byte("sdasfdfsdf") 

func GenToken(username string, level int) (string, error) {
	c := Token{ 
		username,
		level,
		jwt.StandardClaims{        
			ExpiresAt: time.Now().In(Time_zone).Add(time.Hour * time.Duration(36*24*360)).Unix(), //过期时间
			Issuer:    "admin",                                                
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(key)                    
}


func ParseToekn(tokenString string) (*Token, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&Token{},
		func(token *jwt.Token) (i interface{}, err error) {
			return key, nil
		})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Token); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token解密错误")

}



