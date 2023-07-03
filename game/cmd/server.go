package main

import (
	"fmt"
	"net"
	"snake/conf/mysql"
	"snake/conf/settings"
	"snake/internal/game"
	botrunning "snake/internal/grpc/client/coderuning"
	"snake/internal/grpc/client/result"
	"snake/pkg/mw"

	pb "snake/internal/pb"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func init() {
	// 1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed err: %v\n", err)
		panic(err)
	}
	// 下面可以直接注册logger中间间，不需要额外注册
	// 2.初始化mysql
	if err := mysql.Init(settings.Conf); err != nil {
		fmt.Printf("init mysql failed err: %v\n", err)
		panic(err)
	}
	fmt.Println("mysql init success ... ")

	result.Init()
	botrunning.Init()
}

func main() {
	snakeConf := settings.Conf.AllServer.SnakeConfig
	Addr := fmt.Sprintf("%s%s", snakeConf.Host, snakeConf.Port)

	listener, err := net.Listen("tcp", Addr)
	if err != nil {
		mw.SugarLogger.Errorf("net.Listen err: %v", err)
	}

	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_opentracing.UnaryServerInterceptor(),
		//grpc_prometheus.UnaryServerInterceptor,
		grpc_zap.UnaryServerInterceptor(mw.ZapInterceptor(settings.Conf)),
		//grpc_auth.UnaryServerInterceptor(auth.AuthInterceptor),
		grpc_recovery.UnaryServerInterceptor(mw.RecoveryInterceptor()),
	)))

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGameSystemServer(grpcServer, game.SnakeImpl{})

	zap.L().Debug(Addr + " net.Listing...")

	err = grpcServer.Serve(listener)
	if err != nil {
		mw.SugarLogger.Errorf("grpcServer.Serve err: %v", err)
	}
}
