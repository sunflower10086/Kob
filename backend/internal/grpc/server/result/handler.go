package result

import (
	"backend/conf/logger"
	pb "backend/internal/grpc/server/result/pb"

	"golang.org/x/net/context"
)

type ResultServerImpl struct {
	pb.UnimplementedResultServer
}

func (r *ResultServerImpl) Result(ctx context.Context, req *pb.ResultReq) (*pb.ResultResp, error) {
	logger.SugarLogger.Debugf("Result: %v", req)
	return &pb.ResultResp{Message: "success"}, nil
}
