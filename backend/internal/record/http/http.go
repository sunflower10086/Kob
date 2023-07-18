package http

import (
	"backend/internal"
	"backend/internal/record"
	"backend/pkg/ioc"
	"backend/pkg/mw/auth"
	"fmt"

	"github.com/gin-gonic/gin"
)

var handler = &RecordHandler{}

func init() {
	ioc.RegistryGin(handler)
}

type RecordHandler struct {
	svc record.Service
}

func (r *RecordHandler) RouteRegistry(g *gin.Engine) {
	withAuthRouter := g.Group("/api", auth.JwtAuthMiddleware())
	{
		withAuthRouter.GET("/record/getlist/", r.GetRecordList) // 获取用户信息
	}
}

func (r *RecordHandler) Config() error {
	recordService := ioc.GetImpl(internal.RecordServiceName)
	if v, ok := recordService.(record.Service); ok {
		r.svc = v
		return nil
	}

	return fmt.Errorf("%s does not implement the %s Service interface", internal.RecordServiceName, internal.RecordServiceName)
}

func (r *RecordHandler) Name() string {
	return internal.RecordServiceName
}
