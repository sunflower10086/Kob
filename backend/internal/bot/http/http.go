package http

import (
	"backend/internal"
	"backend/internal/bot"
	"backend/pkg/ioc"
	"backend/pkg/mw/auth"
	"fmt"

	"github.com/gin-gonic/gin"
)

var botHandler = &BotHandler{}

// 自注册到Ioc层
func init() {
	ioc.RegistryGin(botHandler)
}

//func NewHostHTTPHandler() *BotHandler {
//	return &BotHandler{}
//}

type BotHandler struct {
	svc bot.Service
}

func (b *BotHandler) Name() string {
	return internal.BotServiceName
}

func (b *BotHandler) Config() error {
	hostService := ioc.GetImpl(internal.BotServiceName)
	if v, ok := hostService.(bot.Service); ok {
		b.svc = v
		return nil
	}

	return fmt.Errorf("%s does not implement the %s Service interface", internal.BotServiceName, internal.BotServiceName)
}

func (b *BotHandler) RouteRegistry(g *gin.Engine) {
	withAuthRouter := g.Group("/api", auth.JwtAuthMiddleware())
	{
		// bot
		botRouter := withAuthRouter.Group("/user/bot")
		botRouter.POST("/add/", b.AddBotHandler)
		botRouter.POST("/remove/", b.RemoveBotHandler)
		botRouter.POST("/update/", b.UpdateBotHandler)
		botRouter.GET("/getlist/", b.GetListBotHandler)
	}
}
