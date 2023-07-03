package main

import (
	"coderunning/conf/settings"
	"coderunning/internal/coderuning"
	"coderunning/internal/coderuning/util"
	"context"
	"fmt"
	"log"
	"net"

	pb "coderunning/internal/pb"

	"google.golang.org/grpc"
)

func init() {
	// 1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed err: %v\n", err)
		panic(err)
	}
}

func main() {
	snakeConf := settings.Conf.AllServer.SnakeConfig
	Addr := fmt.Sprintf("%s%s", snakeConf.Host, snakeConf.Port)

	listener, err := net.Listen("tcp", Addr)
	if err != nil {
		log.Printf("net.Listen err: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCodeRunServer(grpcServer, &coderuning.CodeRunImpl{})

	fmt.Printf(Addr + " net.Listing...")
	ctx := context.Background()
	go util.Run(ctx)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Printf("grpcServer.Serve err: %v", err)
	}
}
