package main

import (
	"fmt"
	"matching/conf/mysql"
	"matching/conf/settings"
	"matching/internal/match"
	"matching/internal/match/logic/matchutil"
	pb "matching/internal/pb/matchingServer"
	"matching/pkg/mw"
	"net"

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

}

func main() {
	Address := fmt.Sprintf("%s%s", settings.Conf.Server.Host, settings.Conf.Server.Port)

	listener, err := net.Listen("tcp", Address)

	if err != nil {
		mw.SugarLogger.Errorf("net.Listen err: %v", err)
	}

	// 添加中间件
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_opentracing.UnaryServerInterceptor(),
		//grpc_prometheus.UnaryServerInterceptor,
		grpc_zap.UnaryServerInterceptor(mw.ZapInterceptor(settings.Conf)),
		//grpc_auth.UnaryServerInterceptor(auth.AuthInterceptor),
		grpc_recovery.UnaryServerInterceptor(mw.RecoveryInterceptor()),
	)))

	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer(opts...)
	// 在gRPC服务器注册我们的服务
	pb.RegisterMatchingSystemServer(grpcServer, &match.MatchingSystemServerImpl{})

	zap.L().Debug(Address + " net.Listing...")
	go matchutil.MatchingPool()

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		mw.SugarLogger.Errorf("grpcServer.Serve err: %v", err)
	}

}
