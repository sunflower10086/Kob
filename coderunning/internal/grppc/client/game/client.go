package game

import (
	"coderunning/conf/settings"
	pb "coderunning/internal/grppc/client/game/pb"
	"context"
	"log"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client pb.GameSystemClient

func Init(conf *settings.AppConfig) {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(conf.AllServer.SnakeConfig.Port, opts...)
	if err != nil {
		log.Printf("snake server net.Connect err: %s\n", err.Error())
	}

	client = pb.NewGameSystemClient(conn)
}

func SetNextStep(ctx context.Context, req *pb.SetNextStepReq) (*pb.SetNextStepResp, error) {
	resp, err := client.SetNextStep(ctx, req)
	if err != nil {
		zap.L().Debug(err.Error())
		return nil, err
	}

	return resp, nil
}
