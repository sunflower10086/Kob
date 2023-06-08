package auth

import (
	"backend/pkg/myerr"
	"backend/pkg/result"
	"backend/pkg/util"
	"github.com/gin-gonic/gin"
)

// JwtAuthMiddleware 登录验证中间件
func JwtAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			result.SendResult(c, result.Fail(myerr.NewMyError("请求的token为空")))
			c.Abort()
			return
		}

		reqToken, err := util.ParseToken(authHeader)
		if err != nil {
			result.SendResult(c, result.Fail(myerr.NewMyError("请求的token验证失败")))
			c.Abort()
			return
		}

		err = util.UpdateTokenExpires(authHeader)
		if err != nil {
			result.SendResult(c, result.Fail(err))
			c.Abort()
			return
		}
		c.Set("user_id", reqToken.UserId)
		c.Next()
	}
}
