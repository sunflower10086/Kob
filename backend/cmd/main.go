package main

import (
	"backend/conf/logger"
	"backend/conf/mysql"
	"backend/conf/redis"
	"backend/conf/settings"
	"backend/internal/routes"
	"backend/pkg/gin_run"
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
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

	// 5.注册路由
	r := routes.Setup()

	// 6.优雅关机
	gin_run.Run(r, viper.GetString("app.name"))
}
