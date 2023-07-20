package client

import (
	"coderunning/conf/settings"
	"coderunning/internal/grppc/client/game"

	"github.com/sunflower10086/Cococola/etcd"
)

func Init() {
	svc := etcd.NewServiceDiscovery([]string{settings.Conf.EtcdConf.Endpoint})
	svc.WatchService("/gRPC/")
	snakeAddr, err := svc.GetService("/gRPC/" + settings.Conf.AllServer.SnakeConfig.Name)
	if err != nil {
		return
	}
	game.Init(snakeAddr)
}
