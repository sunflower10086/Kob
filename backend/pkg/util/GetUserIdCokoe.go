package util

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUserIdCookie(ctx *gin.Context) (int, error) {
	CookieGetId, ok := ctx.Get("user_id")
	if ok == false {
		return -1, errors.New("从cookie中获取userId失败")
	}

	userIdStr := fmt.Sprintf("%v", CookieGetId)
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return -1, err
	}

	return userId, nil
}
