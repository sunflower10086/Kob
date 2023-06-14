package user

import (
	"backend/internal/handlers"
	logicUser "backend/internal/logic/user"
	"backend/pkg/myerr"
	"backend/pkg/result"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LoginHandler 登录接口
// @Summary 登录接口
// @Description 根据用户信息进行登录
// @Tags 用户信息相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query handlers.LoginRequest false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} result.Result
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	var loginParam handlers.LoginRequest

	if err := c.ShouldBind(&loginParam); err != nil {
		zap.L().Debug("param err:", zap.Error(err))
		result.SendResult(c, result.Fail(myerr.ParamErr))
		return
	}

	resp, err := logicUser.LoginService(
		loginParam.UserName,
		loginParam.PassWord,
	)
	if err != nil {
		zap.L().Debug(err.Error())
		result.SendResult(c, result.Fail(err))
		return
	}

	result.SendResult(c, result.Success("success", resp))
}
