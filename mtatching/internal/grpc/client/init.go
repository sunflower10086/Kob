package client

import (
	"matching/conf/settings"
	"matching/internal/grpc/client/result"

	"github.com/sunflower10086/Cococola/etcd"
)

func Init() {
	svc := etcd.NewServiceDiscovery([]string{settings.Conf.EtcdConf.Endpoint})
	svc.WatchService("/gRPC/")

	ResultAddr, err := svc.GetService("/gRPC/" + settings.Conf.AllServer.ResultConfig.Name)
	if err != nil {
		return
	}
	result.Init(ResultAddr)
}
