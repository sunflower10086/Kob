package coderuning

import (
	"coderunning/internal/coderuning/util"
	pb "coderunning/internal/pb"
	"context"
	"strconv"
)

type CodeRunImpl struct {
	pb.UnimplementedCodeRunServer
}

func (c *CodeRunImpl) AddBot(ctx context.Context, req *pb.AddBotReq) (*pb.AddBotResp, error) {
	util.AddBot(req.UserId, req.BotCode, req.Input)
	return &pb.AddBotResp{Message: "success " + strconv.Itoa(int(req.UserId))}, nil
}
