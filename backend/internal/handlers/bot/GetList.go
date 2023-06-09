package bot

import (
	"backend/internal/handlers"
	logicBot "backend/internal/logic/bot"
	"backend/pkg/myerr"
	"backend/pkg/result"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetListBotHandler(c *gin.Context) {
	var getListBotParam handlers.GetListBotReq

	if err := c.ShouldBind(&getListBotParam); err != nil {
		zap.L().Debug(err.Error())
		result.SendResult(c, result.Fail(myerr.ParamErr))
		return
	}

	// 判断请求的参数中userId是否为空
	var UserId string
	if getListBotParam == (handlers.GetListBotReq{}) {
		CookieGetId, _ := c.Get("user_id")
		UserId = fmt.Sprintf("%v", CookieGetId)
	} else {
		UserId = getListBotParam.UserId
	}

	intUserId, err := strconv.Atoi(UserId)
	if err != nil {
		result.SendResult(c, result.Fail(err))
		return
	}

	resp, err := logicBot.GetListService(intUserId)
	if err != nil {
		result.SendResult(c, result.Fail(err))
		return
	}

	result.SendResult(c, result.Success("success", resp))
}
