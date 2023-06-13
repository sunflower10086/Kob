package main

import (
	"backend/conf/logger"
	"backend/conf/mysql"
	"backend/conf/redis"
	"backend/conf/settings"
	"backend/internal/grpc/client"
	"backend/internal/grpc/server"
	"backend/internal/handlers"
	"backend/internal/routes"
	"backend/pkg/gin_run"
	"fmt"

	"go.uber.org/zap"
)

func init() {
	// 1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed err: %v\n", err)
		panic(err)
	}

	// 2.初始化日志
	if err := logger.Init(settings.Conf); err != nil {
		fmt.Printf("init logger failed err: %v\n", err)
		panic(err)
	}
	defer func() {
		zap.L().Sync()
		logger.SugarLogger.Sync()
	}()
	zap.L().Debug("log init success ... ")

	// 3.初始化mysql
	if err := mysql.Init(settings.Conf); err != nil {
		fmt.Printf("init mysql failed err: %v\n", err)
		panic(err)
	}

	// 4.初始化redis
	if err := redis.Init(settings.Conf); err != nil {
		fmt.Printf("init redis failed err: %v\n", err)
		panic(err)
	}
	defer redis.RDB.Close()

	if err := handlers.InitTrans("zh"); err != nil {
		fmt.Printf("init validator failed err: %v\n", err)
	}

	go client.Init()
	go server.Init()
}

// @title Kob
// @version 1.0
// @description 这是一个AI对战平台
// @termsOfService http://www.127.0.0.1:3000/api

// @contact.name 刘钊
// @contact.url http://www.127.0.0.1:3000/api
// @contact.email lz18738377974@163.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:3000
// @BasePath /api
func main() {

	// 5.注册路由
	r := routes.Setup()

	// 6.优雅关机
	gin_run.Run(r, settings.Conf.AllServer.HttpConfig.Port)
}
