package http

import (
	"backend/internal"
	"backend/internal/user"
	"backend/pkg/ioc"
	"backend/pkg/mw/auth"
	"fmt"

	"github.com/gin-gonic/gin"
)

var userHandler = &UserHandler{}

func init() {
	ioc.RegistryGin(userHandler)
}

type UserHandler struct {
	svc user.Service
}

func (u *UserHandler) Name() string {
	return internal.UserServiceName
}

func (u *UserHandler) Config() error {
	hostService := ioc.GetImpl(internal.UserServiceName)
	if v, ok := hostService.(user.Service); ok {
		u.svc = v
		return nil
	}

	return fmt.Errorf("%s does not implement the %s Service interface", internal.UserServiceName, internal.UserServiceName)
}

func (u *UserHandler) RouteRegistry(g *gin.Engine) {
	// noAuth
	noAuthRouter := g.Group("/api")
	{
		noAuthRouter.POST("/user/account/register/", u.RegisterHandler)
		noAuthRouter.POST("/user/account/token/", u.LoginHandler)
	}

	withAuthRouter := g.Group("/api", auth.JwtAuthMiddleware())
	{
		withAuthRouter.GET("/user/account/info/", u.GetInfoHandler) // 获取用户信息
	}
}
