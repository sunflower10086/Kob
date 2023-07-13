package protocol

import (
	"backend/conf/settings"
	"backend/internal/routes"
	"backend/pkg/ioc"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func NewHTTPService() *HTTPService {
	g := gin.Default()
	httpConf := settings.Conf

	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              fmt.Sprintf("%s:%s", httpConf.AllServer.HttpConfig.Host, httpConf.AllServer.HttpConfig.Port),
		Handler:           g,
	}

	return &HTTPService{
		Service: server,
		L:       log.New(os.Stderr, "[Host] ", log.Ldate|log.Ltime|log.Lshortfile),
		Conf:    httpConf,
		engin:   g,
		Addr:    fmt.Sprintf("%s:%s", httpConf.AllServer.HttpConfig.Host, httpConf.AllServer.HttpConfig.Port),
	}
}

type HTTPService struct {
	Service *http.Server
	L       *log.Logger
	Conf    *settings.AppConfig
	engin   *gin.Engine
	Addr    string
}

func (h *HTTPService) Start() error {
	routes.OtherSetup(h.engin)

	if err := ioc.InitGinHandler(h.engin); err != nil {
		fmt.Println(err)
		return err
	}

	// 已加载的app的日志信息打印
	handlers := ioc.LoadedGinHandler()
	h.L.Println("loaded handler has", handlers)

	// 监听端口
	h.L.Printf("%s running in %s \n", h.Conf.AllServer.HttpConfig.Name, h.Addr)
	if err := h.Service.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			h.L.Fatalf("listen: %s\n", err)
			return nil
		} else {
			return fmt.Errorf("start service err: %s", err.Error())
		}
	}
	return nil
}

func (h *HTTPService) Stop() {
	h.L.Println("start graceful shutdown")
	// 开始关闭
	timeOut := 2
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := h.Service.Shutdown(ctx); err != nil {
		h.L.Fatalf("%s Shutdown err: %v\n", h.Addr, err)
	}
	// 关闭超时
	select {
	case <-ctx.Done():
		h.L.Printf("timeout of %d seconds.", timeOut)
	}

	// 正常关闭
	h.L.Printf("%s exiting", h.Conf.AllServer.HttpConfig.Name)
}
