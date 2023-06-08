package bot

import (
	"backend/internal/handlers"
	logicBot "backend/internal/logic/bot"
	"backend/pkg/myerr"
	"backend/pkg/result"
	"backend/pkg/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RemoveBotHandler(c *gin.Context) {
	var removeBotParam handlers.RemoveBotReq

	if err := c.ShouldBind(&removeBotParam); err != nil {
		result.SendResult(c, result.Fail(myerr.ParamErr))
		return
	}

	intBotId, err := strconv.Atoi(removeBotParam.BotId)
	if err != nil {
		result.SendResult(c, result.Fail(err))
		return
	}

	userId, err := util.GetUserIdCookie(c)
	if err != nil {
		result.SendResult(c, result.Fail(err))
		return
	}

	resp, err := logicBot.RemoveService(
		intBotId,
		userId,
	)
	if err != nil {
		result.SendResult(c, result.Fail(err))
		return
	}

	result.SendResult(c, result.Success(resp))
}
