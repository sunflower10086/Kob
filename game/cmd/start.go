package cmd

import (
	"fmt"
	"net"
	"snake/conf/mysql"
	"snake/conf/settings"
	"snake/internal/game"
	botrunning "snake/internal/grpc/client/coderuning"
	"snake/internal/grpc/client/result"
	pb "snake/internal/pb"
	"snake/pkg/mw"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/spf13/cobra"
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

		result.Init()
		botrunning.Init()

		gameGrpcStart()
		return nil
	},
}

func init() {

	StartCmd.PersistentFlags().StringVarP(&configFile, "config", "f", "./etc/config.yaml", "demo config file")
	RootCmd.AddCommand(StartCmd)
}

func gameGrpcStart() {
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
