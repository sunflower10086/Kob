package client

import (
	"snake/conf/settings"
	code "snake/internal/grpc/client/coderuning"
	"snake/internal/grpc/client/result"

	"github.com/sunflower10086/Cococola/etcd"
)

func Init() {
	svc := etcd.NewServiceDiscovery([]string{settings.Conf.EtcdConf.Endpoint})
	svc.WatchService("/gRPC/")

	BotRunningAddr, err := svc.GetService("/gRPC/" + settings.Conf.AllServer.BotRunningConfig.Name)
	if err != nil {
		return
	}
	code.Init(BotRunningAddr)

	resultConfig, err := svc.GetService("/gRPC/" + settings.Conf.AllServer.ResultConfig.Name)
	if err != nil {
		return
	}
	result.Init(resultConfig)

}
