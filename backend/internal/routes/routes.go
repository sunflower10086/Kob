package routes

import (
	"backend/conf"
	"backend/conf/logger"
	"backend/internal/consumer"
	"backend/pkg/result"
	"backend/pkg/util"
	"net/http"
	"strconv"

	_ "backend/docs" // 千万不要忘了导入把你上一步生成的docs

	"github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func OtherSetup(r *gin.Engine) {

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

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusOK, "no route")
	})

	r.NoMethod(func(c *gin.Context) {
		c.String(http.StatusOK, "no method")
	})
}
