package util

import (
	"blog/pkg/setting"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username           string `json:"username"`
	Password           string `json:"password"`
	jwt.StandardClaims        //内置结构体  它包含一些标准的 JWT 声明信息，例如过期时间、签发时间等。
}

// 生成token的go函数  传入用户名和密码，返回一个token和错误
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour) //设置当前时间加上三个小时
	//1.声明需要用户名，密码，jwt声明标准
	claims := Claims{
		username,
		password,
		//jwt标准中定义的一组申明，
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //指定token的过期时间，unix时间戳
			Issuer:    "gin-blog",        //指定token签发人，用于验证token合法性
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	//2.jwt.NewWithClaims 函数创建一个 Token 对象,用于进行签名。使用 HMAC-SHA256 签名算法进行签名
	fmt.Println(tokenClaims)
	token, err := tokenClaims.SignedString(jwtSecret)
	//3.tokenClaims.SignedString 方法对 Token 进行签名  jwtSecret密钥
	return token, err
}

// 解码token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
