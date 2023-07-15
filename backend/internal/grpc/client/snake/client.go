package snake

import (
	"backend/conf/logger"
	"backend/conf/settings"
	"backend/internal/grpc/client/snake/util"
	shape "backend/pkg/share_space"
	"context"
	"fmt"

	pb "backend/internal/grpc/client/snake/pb"

	share "backend/pkg/share_space"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client  pb.GameSystemClient
	SnakeMd *WithSnake
	Space   *shape.Space
)

func Init(conf *settings.AppConfig) {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(conf.AllServer.SnakeConfig.GetAddr(), opts...)
	if err != nil {
		zap.L().Error("snake server net.Connect err: ", zap.Error(err))
	}

	client = pb.NewGameSystemClient(conn)

	Space = share.NewSpace()

	SnakeMd = &WithSnake{
		Msg: make(chan shape.Pair),
	}

	var forGameSystem util.CommGame

	forGameSystem = SnakeMd

	// 向game_system发消息
	go func() {
		logger.SugarLogger.Debug("Send Message")
		err := forGameSystem.Send()
		if err != nil {
			logger.SugarLogger.Debugf("Send Message err: %v", err)
			return
		}
	}()
}

func StartGame(ctx context.Context, req *pb.StartGameReq) (*pb.StartGameResp, error) {
	resp, err := client.StartGame(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func SetNextStep(ctx context.Context, req *pb.SetNextStepReq) (*pb.SetNextStepResp, error) {
	resp, err := client.SetNextStep(ctx, req)
	if err != nil {
		zap.L().Debug(err.Error())
		fmt.Println(err.Error())
		return nil, err
	}

	return resp, nil
}
