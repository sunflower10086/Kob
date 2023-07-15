package match

import (
	"backend/conf/settings"
	pb "backend/internal/grpc/client/match/pb"

	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var matchSysClient pb.MatchingSystemClient

func Init(conf *settings.AppConfig) {
	// 连接服务器
	var opts []grpc.DialOption

	// 明文传输，不做认证
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(conf.AllServer.MatchConfig.GetAddr(), opts...)
	if err != nil {
		zap.L().Error("match server net.Connect err: ", zap.Error(err))
	}

	// 建立gRPC连接
	matchSysClient = pb.NewMatchingSystemClient(conn)
}

func AddUser(ctx context.Context, user *pb.User) (*pb.Response, error) {

	resp, err := matchSysClient.AddUser(ctx, user)
	if err != nil {
		return nil, err
	}
	zap.L().Debug(resp.Message)
	return resp, nil
}

func RemoveUser(ctx context.Context, user *pb.User) (*pb.Response, error) {
	resp, err := matchSysClient.Remove(ctx, user)
	if err != nil {
		return nil, err
	}
	zap.L().Debug(resp.Message)
	return resp, nil
}
