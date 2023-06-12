package client

import (
	"context"
	"log"
	"matching/conf/settings"
	pb "matching/internal/grpc/client/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var resultClient pb.ResultClient

func InitResult() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(settings.Conf.AllServer.ResultConfig.Port, opts...)
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}

	resultClient = pb.NewResultClient(conn)
}

func Result(ctx context.Context, result *pb.ResultReq) (*pb.ResultResp, error) {
	log.Println(result)
	resp, err := resultClient.Result(ctx, result)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
