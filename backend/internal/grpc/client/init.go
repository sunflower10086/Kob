package client

import (
	"backend/conf/settings"
	"backend/internal/grpc/client/match"
	"backend/internal/grpc/client/snake"
	"fmt"

	"github.com/sunflower10086/Cococola/etcd"
	"go.uber.org/zap"
)

func Init() {
	fmt.Println("client init1")
	svc := etcd.NewServiceDiscovery([]string{settings.Conf.EtcdConf.Endpoint})
	svc.WatchService("/gRPC/")
	matchAddr, err := svc.GetService("/gRPC/" + settings.Conf.AllServer.MatchConfig.Name)
	if err != nil {
		zap.L().Debug(err.Error())
		return
	}
	match.Init(matchAddr)

	fmt.Println("client init2")
	snakeAddr, err := svc.GetService("/gRPC/" + settings.Conf.AllServer.SnakeConfig.Name)
	if err != nil {
		zap.L().Debug(err.Error())
		return
	}
	snake.Init(snakeAddr)

}
