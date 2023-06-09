package main

import (
	"fmt"
	"matching/conf/logger"
	"matching/conf/mysql"
	"matching/conf/settings"
	"matching/internal/match"
	"matching/internal/match/logic/matchutil"
	pb "matching/internal/pb/matchingServer"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
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
	zap.L().Debug("mysql init success ... ")

}

func main() {
	Address := fmt.Sprintf("%s%s", settings.Conf.Host, settings.Conf.Port)

	listener, err := net.Listen("tcp", Address)

	if err != nil {
		logger.SugarLogger.Errorf("net.Listen err: %v", err)
	}
	zap.L().Debug(Address + " net.Listing...")
	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	pb.RegisterMatchingSystemServer(grpcServer, &match.MatchingSystemServerImpl{})

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		logger.SugarLogger.Errorf("grpcServer.Serve err: %v", err)
	}

	go matchutil.MatchingPool()
}
