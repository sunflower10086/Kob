package cmd

import (
	"backend/conf/logger"
	"backend/conf/mysql"
	"backend/conf/redis"
	"backend/conf/settings"
	"backend/internal/grpc/client"
	"backend/internal/grpc/server"
	"backend/pkg/ioc"
	"backend/protocol"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "backend/internal/all"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	configFile string
)

var StartCmd = &cobra.Command{
	Use:     "start",
	Long:    "backend后端",
	Short:   "backend后端",
	Example: "backend后端 commands",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 1.加载配置
		if err := settings.Init(); err != nil {
			fmt.Printf("init settings failed err: %v\n", err)
			return err
		}

		// 2.初始化日志
		if err := logger.Init(settings.Conf); err != nil {
			fmt.Printf("init logger failed err: %v\n", err)
			return err
		}
		defer func() {
			zap.L().Sync()
			logger.SugarLogger.Sync()
		}()
		zap.L().Debug("log init success ... ")

		// 3.初始化mysql
		if err := mysql.Init(settings.Conf); err != nil {
			fmt.Printf("init mysql failed err: %v\n", err)
			return err
		}

		// 4.初始化redis
		if err := redis.Init(settings.Conf); err != nil {
			fmt.Printf("init redis failed err: %v\n", err)
			return err
		}
		defer redis.RDB.Close()

		go client.Init()
		go server.Init()

		ioc.InitImpl()

		master := newMaster()

		// 相当于监听一下 kill -2 和 kill -9
		quit := make(chan os.Signal)
		// kill (no param) default send syscanll.SIGTERM
		// kill -2 is syscall.SIGINT (Ctrl + C)
		// kill -9 is syscall.SIGKILL
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		go master.WaitStop(quit)

		return master.Start()
	},
}

func newMaster() *master {
	return &master{
		http: protocol.NewHTTPService(),
	}
}

type master struct {
	http *protocol.HTTPService
}

func (m *master) Start() error {
	if err := m.http.Start(); err != nil {
		return err
	}
	return nil
}

func (m *master) Stop() {
	log.Printf("Shutdown %s ...\n", m.http.Conf.AllServer.HttpConfig.Name)
}

func (m *master) WaitStop(quit <-chan os.Signal) {
	for v := range quit {
		switch v {
		default:
			m.http.L.Printf("received signal: %s", v)
			m.http.Stop()
		}
	}
}

func init() {

	StartCmd.PersistentFlags().StringVarP(&configFile, "config", "f", "./etc/demo.toml", "demo config file")
	RootCmd.AddCommand(StartCmd)
}
