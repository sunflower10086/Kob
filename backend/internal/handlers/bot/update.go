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

func UpdateBotHaandler(c *gin.Context) {
	var updateBotParam handlers.UpdateBotReq

	if err := c.ShouldBind(&updateBotParam); err != nil {
		result.SendResult(c, result.Fail(myerr.ParamErr))
		return
	}

	intBotId, err := strconv.Atoi(updateBotParam.BotId)
	if err != nil {
		result.SendResult(c, result.Fail(err))
		return
	}

	userId, err := util.GetUserIdCookie(c)
	if err != nil {
		result.SendResult(c, result.Fail(err))
		return
	}

	resp, err := logicBot.UpdateBotService(
		intBotId,
		userId,
		updateBotParam.Title,
		updateBotParam.Code,
		updateBotParam.Description,
	)
	if err != nil {
		result.SendResult(c, result.Fail(err))
		return
	}

	result.SendResult(c, result.Success("success", resp))
}
