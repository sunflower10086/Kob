package result

import (
	"context"
	"log"
	"matching/internal/grpc/client/result/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var resultClient result.ResultClient

func Init(endpoint string) {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}

	resultClient = result.NewResultClient(conn)
}

func Result(ctx context.Context, result *result.ResultReq) (*result.ResultResp, error) {
	log.Println(result)
	resp, err := resultClient.Result(ctx, result)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
