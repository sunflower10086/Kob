package snake

import (
	"backend/conf/logger"
	"backend/conf/settings"
	"backend/internal/grpc/client/snake/util"
	shape "backend/pkg/share_space"
	"context"

	pb "backend/internal/grpc/client/snake/pb"

	share "backend/pkg/share_space"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	Client  pb.GameSystemClient
	SnakeMd *WithSnake
	Space   *shape.Space
)

func Init(conf *settings.AppConfig) {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(conf.AllServer.SnakeConfig.Port, opts...)
	if err != nil {
		zap.L().Error("snake server net.Connect err: ", zap.Error(err))
	}

	Client = pb.NewGameSystemClient(conn)

	Space = share.NewSpace()

	SnakeMd = &WithSnake{
		Msg:        make(chan shape.Pair),
		GameClient: Client,
	}

	var forGameSystem util.CommGame

	forGameSystem = SnakeMd
	// 接收从game_system发来的消息
	go func() {
		err := forGameSystem.Receive()
		if err != nil {
			logger.SugarLogger.Debugf("Receive Message err: %v", err)
			return
		}
	}()

	// 向game_system发消息
	go func() {
		err := forGameSystem.Send()
		if err != nil {
			logger.SugarLogger.Debugf("Send Message err: %v", err)
			return
		}
	}()
}

func StartGame(ctx context.Context, req *pb.StartGameReq) (*pb.StartGameResp, error) {
	resp, err := Client.StartGame(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
