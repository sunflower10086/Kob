package routes

import (
	"backend/conf"
	"backend/conf/logger"
	"backend/internal/consumer"
	"backend/internal/handlers/bot"
	"backend/internal/handlers/user"
	"backend/pkg/mw/auth"
	"backend/pkg/result"
	"backend/pkg/util"
	"net/http"
	"strconv"

	_ "backend/docs" // 千万不要忘了导入把你上一步生成的docs

	"github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	r.Use(logger.GinLog(), gin.Recovery(), conf.Cors())

	// 把webSocket的注册中心运行起来
	var hub = consumer.NewHub()

	go hub.Run()
	r.GET("/ws/:token", func(ctx *gin.Context) {
		token := ctx.Param("token")
		parseToken, err := util.ParseToken(token)
		if err != nil {
			result.SendResult(ctx, result.Fail(err))
			return
		}
		userId := parseToken.UserId
		consumer.WsHandler(ctx, hub, strconv.Itoa(userId))
	})

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	// noAuth
	noAuthRouter := r.Group("/api")
	{
		noAuthRouter.POST("/user/account/register/", user.RegisterHandler)
		noAuthRouter.POST("/user/account/token/", user.LoginHandler)
	}

	//withAuth
	withAuthRouter := r.Group("/api", auth.JwtAuthMiddleware())
	{
		withAuthRouter.GET("/user/account/info/", user.GetInfoHandler) // 获取用户信息

		// bot
		botRouter := withAuthRouter.Group("/user/bot")
		botRouter.POST("/add/", bot.AddBotHandler)
		botRouter.POST("/remove/", bot.RemoveBotHandler)
		botRouter.POST("/update/", bot.UpdateBotHaandler)
		botRouter.GET("/getlist/", bot.GetListBotHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusOK, "no route")
	})

	r.NoMethod(func(c *gin.Context) {
		c.String(http.StatusOK, "no method")
	})

	return r
}
