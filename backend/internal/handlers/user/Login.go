package user

import (
	"backend/internal/handlers"
	logicUser "backend/internal/logic/user"
	"backend/pkg/myerr"
	"backend/pkg/result"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var loginParam handlers.LoginRequest

	if err := c.ShouldBind(&loginParam); err != nil {
		result.SendResult(c, result.Fail(myerr.ParamErr))
		return
	}

	resp, err := logicUser.LoginService(
		loginParam.UserName,
		loginParam.PassWord,
	)
	if err != nil {
		result.SendResult(c, result.Fail(err))
		return
	}

	result.SendResult(c, result.Success(resp))
}
