package cmd

import (
	"coderunning/conf/settings"
	"coderunning/internal/coderuning"
	"coderunning/internal/coderuning/util"
	"coderunning/internal/grppc/client/game"
	pb "coderunning/internal/pb"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
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
			panic(err)
		}

		game.Init(settings.Conf)

		botRunningConf := settings.Conf.AllServer.BotRunningConfig
		Addr := fmt.Sprintf("%s%s", botRunningConf.Host, botRunningConf.Port)

		listener, err := net.Listen("tcp", Addr)
		if err != nil {
			log.Printf("net.Listen err: %v", err)
		}

		var opts []grpc.ServerOption

		grpcServer := grpc.NewServer(opts...)
		pb.RegisterCodeRunServer(grpcServer, &coderuning.CodeRunImpl{})

		fmt.Printf(Addr + " net.Listing...\n")
		ctx := context.Background()
		go util.Run(ctx)

		err = grpcServer.Serve(listener)
		if err != nil {
			log.Printf("grpcServer.Serve err: %v", err)
		}
		return nil
	},
}

func init() {

	StartCmd.PersistentFlags().StringVarP(&configFile, "config", "f", "./etc/demo.toml", "demo config file")
	RootCmd.AddCommand(StartCmd)
}
