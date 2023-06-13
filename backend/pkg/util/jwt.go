package util

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type MyClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

var Secret = []byte("secret")

const TokenExpireDuration = time.Hour * 24

// CreateToken 生成Token
func CreateToken(userId int) (string, error) {
	t := MyClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			// 签名时间
			IssuedAt: time.Now().Unix(),
			Issuer:   "kob",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, t)
	// c.Set("token", token)
	return token.SignedString(Secret)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*MyClaims, error) {
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// UpdateTokenExpires 更新token的有效时间
func UpdateTokenExpires(token string) error {
	claims, err := ParseToken(token)
	if err != nil {
		return err
	}

	// 更新token的有效时间
	if time.Now().Unix() < claims.ExpiresAt {
		claims.ExpiresAt = time.Now().Add(TokenExpireDuration).Unix()
		//// 签名时间
		claims.IssuedAt = time.Now().Unix()
	}

	return nil
}

// GenToken ⽣生成access token 和 refresh token
func GenToken(userID int) (aToken, rToken string, err error) {
	// 创建⼀一个我们⾃自⼰己的声明
	c := MyClaims{
		userID, // ⾃自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "bluebell",                                 // 签发⼈人
		},
	}
	// 加密并获得完整的编码后的字符串串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256,
		c).SignedString(Secret)
	// refresh token 不不需要存任何⾃自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(), // 过期时间
		Issuer:    "bluebell",                              // 签发⼈人
	}).SignedString(Secret)
	// 使⽤用指定的secret签名并获得完整的编码后的字符串串token
	return
}
