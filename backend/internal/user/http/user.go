package http

import (
	"backend/internal/user"
	"backend/pkg/myerr"
	"backend/pkg/result"
	"backend/pkg/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (u *UserHandler) GetInfoHandler(c *gin.Context) {
	var getUserInfoParam user.GetUserInfoRequest

	if err := c.ShouldBind(&getUserInfoParam); err != nil {
		result.SendResult(c, result.Fail(myerr.ParamErr))
		return
	}

	var UserId int
	var err error
	if getUserInfoParam == (user.GetUserInfoRequest{}) {
		UserId, err = util.GetUserIdCookie(c)
	} else {
		UserId, err = strconv.Atoi(getUserInfoParam.UserId)
	}

	if err != nil {
		result.SendResult(c, result.Fail(err))
		return
	}

	resp, err := u.svc.GetInfoService(
		c,
		&user.GetUserInfoRequest{
			UserId: strconv.Itoa(UserId),
		})
	if err != nil {
		result.SendResult(c, result.Fail(err))
		return
	}

	result.SendResult(c, result.Success("success", resp))
}

func (u *UserHandler) LoginHandler(c *gin.Context) {
	var loginParam user.LoginRequest

	if err := c.ShouldBind(&loginParam); err != nil {
		zap.L().Debug("param err:", zap.Error(err))
		result.SendResult(c, result.Fail(myerr.ParamErr))
		return
	}

	resp, err := u.svc.LoginService(
		c,
		&user.LoginRequest{
			UserName: loginParam.UserName,
			PassWord: loginParam.PassWord,
		},
	)
	if err != nil {
		zap.L().Debug(err.Error())
		result.SendResult(c, result.Fail(err))
		return
	}

	result.SendResult(c, result.Success("success", resp))
}

func (u *UserHandler) RegisterHandler(c *gin.Context) {
	registerParam := user.RegisterRequest{}

	if err := c.ShouldBind(&registerParam); err != nil {
		result.SendResult(c, result.Fail(myerr.ParamErr))
		return
	}

	// TODO: 写逻辑
	resp, err := u.svc.RegisterService(
		c,
		&user.RegisterRequest{
			UserName:          registerParam.UserName,
			PassWord:          registerParam.PassWord,
			ConfirmedPassword: registerParam.ConfirmedPassword,
		},
	)
	if err != nil {
		zap.L().Error(err.Error())
		result.SendResult(c, result.Fail(err))
		return
	}

	result.SendResult(c, result.Success("success", resp))
}
