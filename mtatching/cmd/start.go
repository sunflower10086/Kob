package cmd

import (
	"fmt"
	"matching/conf/mysql"
	"matching/conf/settings"
	"matching/internal/grpc/client"
	"matching/internal/match"
	"matching/internal/match/logic/matchutil"
	pb "matching/internal/pb/matchingServer"
	"matching/pkg/mw"
	"net"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/spf13/cobra"
	"github.com/sunflower10086/Cococola/etcd"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	configFile string
)

var StartCmd = &cobra.Command{
	Use:     "start",
	Long:    "code running 微服务",
	Short:   "code running 微服务",
	Example: "code running 微服务 commands",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 1.加载配置
		if err := settings.Init(); err != nil {
			fmt.Printf("init settings failed err: %v\n", err)
			return err
		}
		// 下面可以直接注册logger中间间，不需要额外注册

		// 2.初始化mysql
		if err := mysql.Init(settings.Conf); err != nil {
			fmt.Printf("init mysql failed err: %v\n", err)
			return err
		}
		fmt.Println("mysql init success ... ")

		go client.Init()

		matchGrpcStart()
		return nil
	},
}

func init() {

	StartCmd.PersistentFlags().StringVarP(&configFile, "config", "f", "./etc/config.yaml", "demo config file")
	RootCmd.AddCommand(StartCmd)
}

func matchGrpcStart() {
	matchConf := settings.Conf.AllServer.MatchConfig
	Address := matchConf.GetAddr()

	listener, err := net.Listen("tcp", Address)

	if err != nil {
		mw.SugarLogger.Errorf("net.Listen err: %v", err)
	}

	// 添加中间件
	var opts []grpc.ServerOption
	opts = append(opts, grpc.ChainUnaryInterceptor(
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_opentracing.UnaryServerInterceptor(),
		//grpc_prometheus.UnaryServerInterceptor,
		grpc_zap.UnaryServerInterceptor(mw.ZapInterceptor(settings.Conf)),
		//grpc_auth.UnaryServerInterceptor(auth.AuthInterceptor),
		grpc_recovery.UnaryServerInterceptor(mw.RecoveryInterceptor()),
	))

	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer(opts...)
	// 在gRPC服务器注册我们的服务
	pb.RegisterMatchingSystemServer(grpcServer, &match.MatchingSystemServerImpl{})

	zap.L().Debug(Address + " net.Listing...")
	go matchutil.MatchingPool()

	svc, err := etcd.NewServiceRegister([]string{settings.Conf.EtcdConf.Endpoint},
		"/gRPC/"+matchConf.Name,
		matchConf.GetAddr(),
		3,
	)
	if err != nil {
		return
	}
	go svc.ListenLeaseRespChan()

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		mw.SugarLogger.Errorf("grpcServer.Serve err: %v", err)
	}
}
