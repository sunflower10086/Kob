package result

import (
	"context"
	"log"
	pb "snake/internal/grpc/client/result/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var resultClient pb.ResultClient

func Init(endpoint string) {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(endpoint, opts...)
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
