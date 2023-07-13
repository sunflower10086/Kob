package http

import (
	"backend/internal/bot"
	"backend/pkg/myerr"
	"backend/pkg/result"
	"backend/pkg/util"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *BotHandler) AddBotHandler(c *gin.Context) {
	var addBotParam bot.AddBotReq

	// 参数校验
	if err := c.ShouldBind(&addBotParam); err != nil {
		zap.L().Debug(err.Error())
		result.SendResult(c, result.Fail(myerr.ParamErr))
		return
	}
	// 从cookie中获取userId
	userId, err := util.GetUserIdCookie(c)
	if err != nil {
		result.SendResult(c, result.Fail(err))
		return
	}

	resp, err := h.svc.AddBot(
		c,
		&bot.AddBotReq{
			Title:       addBotParam.Title,
			Code:        addBotParam.Code,
			Description: addBotParam.Description,
			UserId:      userId,
		},
	)
	if err != nil {
		result.SendResult(c, result.Fail(err))
		return
	}

	result.SendResult(c, result.Success("success", resp))
}

func (h *BotHandler) GetListBotHandler(c *gin.Context) {
	var getListBotParam bot.GetListBotReq

	if err := c.ShouldBind(&getListBotParam); err != nil {
		zap.L().Debug(err.Error())
		result.SendResult(c, result.Fail(myerr.ParamErr))
		return
	}

	// 判断请求的参数中userId是否为空
	var UserId string
	if getListBotParam == (bot.GetListBotReq{}) {
		CookieGetId, _ := c.Get("user_id")
		UserId = fmt.Sprintf("%v", CookieGetId)
	} else {
		UserId = getListBotParam.UserId
	}

	resp, err := h.svc.GetListBot(
		c,
		&bot.GetListBotReq{
			UserId: UserId,
		})
	if err != nil {
		result.SendResult(c, result.Fail(err))
		return
	}

	result.SendResult(c, result.Success("success", resp))
}

func (h *BotHandler) RemoveBotHandler(c *gin.Context) {
	var removeBotParam bot.DeleteBotReq

	if err := c.ShouldBind(&removeBotParam); err != nil {
		result.SendResult(c, result.Fail(myerr.ParamErr))
		return
	}

	userId, err := util.GetUserIdCookie(c)
	if err != nil {
		result.SendResult(c, result.Fail(err))
		return
	}

	resp, err := h.svc.DeleteBot(
		c,
		&bot.DeleteBotReq{
			UserId: userId,
			BotId:  removeBotParam.BotId,
		},
	)
	if err != nil {
		result.SendResult(c, result.Fail(err))
		return
	}

	result.SendResult(c, result.Success("success", resp))
}

func (h *BotHandler) UpdateBotHandler(c *gin.Context) {
	var updateBotParam bot.UpdateBotReq

	if err := c.ShouldBind(&updateBotParam); err != nil {
		result.SendResult(c, result.Fail(myerr.ParamErr))
		return
	}

	userId, err := util.GetUserIdCookie(c)
	if err != nil {
		result.SendResult(c, result.Fail(err))
		return
	}

	resp, err := h.svc.UpdateBot(
		c,
		&bot.UpdateBotReq{
			UserId:      strconv.Itoa(userId),
			BotId:       updateBotParam.BotId,
			Title:       updateBotParam.Title,
			Code:        updateBotParam.Code,
			Description: updateBotParam.Description,
		},
	)
	if err != nil {
		result.SendResult(c, result.Fail(err))
		return
	}

	result.SendResult(c, result.Success("success", resp))
}
