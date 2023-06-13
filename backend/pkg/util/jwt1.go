package util

// Access Token 和 Refresh Token模式
//import "github.com/golang-jwt/jwt"
//
//// ParseToken1 解析JWT
//func ParseToken1(tokenString string) (claims *MyClaims, err error) {
//	// 解析token
//	var token *jwt.Token
//	claims = new(MyClaims)
//	token, err = jwt.ParseWithClaims(tokenString, claims, keyFunc)
//	if err != nil {
//		return
//	}
//	if !token.Valid { // 校验token
//		err = errors.New("invalid token")
//	}
//	return
//}
//
//// RefreshToken1 刷新AccessToken
//func RefreshToken1(aToken, rToken string) (newAToken, newRToken string, err error) {
//	// refresh token⽆无效直接返回
//	if _, err = jwt.Parse(rToken, keyFunc); err != nil {
//		return
//	}
//	// 从旧access token中解析出claims数据
//	var claims MyClaims
//	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
//	v, _ := err.(*jwt.ValidationError)
//	// 当access token是过期错误 并且 refresh token没有过期时就创建⼀一个新的access token
//	if v.Errors == jwt.ValidationErrorExpired {
//		return GenToken(claims.UserId)
//	}
//	return
//}
