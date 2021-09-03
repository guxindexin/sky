package jwtauth

import (
	"errors"
	"sky/app/system/models"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type Claims struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userInfo *models.User) (string, error) {
	// 创建一个我们自己的声明
	c := Claims{
		userInfo.Username,
		userInfo.IsAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(viper.GetInt("jwt.expires"))).Unix(), // 过期时间
			Issuer:    viper.GetString("jwt.issuer"),                                                   // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString([]byte(viper.GetString("jwt.secret")))
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*Claims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(viper.GetString("jwt.secret")), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
