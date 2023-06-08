package bot

import (
	"backend/internal/handlers"
	logicBot "backend/internal/logic/bot"
	"backend/pkg/myerr"
	"backend/pkg/result"
	"backend/pkg/util"

	"github.com/gin-gonic/gin"
)

func AddBotHandler(c *gin.Context) {
	var addBotParam handlers.AddBotReq

	if err := c.ShouldBind(&addBotParam); err != nil {
		result.SendResult(c, result.Fail(myerr.ParamErr))
		return
	}
	userId, err := util.GetUserIdCookie(c)
	if err != nil {
		result.SendResult(c, result.Fail(err))
		return
	}

	resp, err := logicBot.AddBotService(
		addBotParam.Title,
		addBotParam.Code,
		addBotParam.Description,
		userId,
	)
	if err != nil {
		result.SendResult(c, result.Fail(err))
		return
	}

	result.SendResult(c, result.Success(resp))
}
