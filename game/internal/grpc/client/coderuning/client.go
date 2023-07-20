package result

import (
	"context"
	"fmt"
	"log"
	"snake/conf/settings"
	pb "snake/internal/grpc/client/coderuning/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var botRunningClient pb.CodeRunClient

func Init(endpoint string) {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	fmt.Println(settings.Conf.AllServer.BotRunningConfig)
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}

	botRunningClient = pb.NewCodeRunClient(conn)
}

func AddBot(ctx context.Context, req *pb.AddBotReq) (*pb.AddBotResp, error) {
	resp, err := botRunningClient.AddBot(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return resp, nil
}
