package user

import (
	"backend/internal/handlers"
	logicUser "backend/internal/logic/user"
	"backend/pkg/myerr"
	"backend/pkg/result"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	registerParam := handlers.RegisterRequest{}

	if err := c.ShouldBind(&registerParam); err != nil {
		result.SendResult(c, result.Fail(myerr.ParamErr))
		return
	}

	// TODO: 写逻辑
	resp, err := logicUser.RegisterService(
		registerParam.UserName,
		registerParam.PassWord,
		registerParam.ConfirmedPassword,
	)
	if err != nil {
		result.SendResult(c, result.Fail(myerr.ParamErr))
		return
	}

	result.SendResult(c, result.Success(resp))
}
