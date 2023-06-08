package user

import (
	"backend/internal/handlers"
	logicUser "backend/internal/logic/user"
	"backend/pkg/myerr"
	"backend/pkg/result"
	"backend/pkg/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetInfoHandler(c *gin.Context) {
	var getUserInfoParam handlers.GetUserInfoRequest

	if err := c.ShouldBind(&getUserInfoParam); err != nil {
		result.SendResult(c, result.Fail(myerr.ParamErr))
		return
	}

	var UserId int
	var err error
	if getUserInfoParam == (handlers.GetUserInfoRequest{}) {
		UserId, err = util.GetUserIdCookie(c)
	} else {
		UserId, err = strconv.Atoi(getUserInfoParam.UserId)
	}

	if err != nil {
		result.SendResult(c, result.Fail(err))
		return
	}

	resp, err := logicUser.GetInfoService(UserId)
	if err != nil {
		result.SendResult(c, result.Fail(err))
		return
	}

	result.SendResult(c, result.Success(resp))
}
