package http

import (
	"backend/internal"
	"backend/internal/rank"
	"backend/pkg/ioc"
	"backend/pkg/mw/auth"
	"fmt"

	"github.com/gin-gonic/gin"
)

var handler = &RankHandler{}

func init() {
	ioc.RegistryGin(handler)
}

type RankHandler struct {
	svc rank.Service
}

func (r *RankHandler) RouteRegistry(g *gin.Engine) {
	withAuthRouter := g.Group("/api", auth.JwtAuthMiddleware())
	{
		withAuthRouter.GET("/rank/getlist/", r.GetRankList) // 获取用户信息
	}
}

func (r *RankHandler) Config() error {
	rankHandler := ioc.GetImpl(internal.RankServiceName)
	if service, ok := rankHandler.(rank.Service); ok {
		r.svc = service
		return nil
	}

	return fmt.Errorf("%s does not implement the %s Service interface", internal.RankServiceName, internal.RankServiceName)

}

func (r *RankHandler) Name() string {
	return internal.RankServiceName
}
