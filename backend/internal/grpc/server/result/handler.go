package result

import (
	"backend/conf/logger"
	"backend/internal/grpc/client/snake"
	snakePb "backend/internal/grpc/client/snake/pb"
	resultPb "backend/internal/grpc/server/result/pb"

	"go.uber.org/zap"
	"golang.org/x/net/context"
)

const (
	startGameType int32 = iota
	gameResultType
)

type ResultServerImpl struct {
	resultPb.UnimplementedResultServer
}

func (r *ResultServerImpl) Result(ctx context.Context, req *resultPb.ResultReq) (*resultPb.ResultResp, error) {
	logger.SugarLogger.Debugf("Result: %v", req)
	switch req.GetEventType() {
	case startGameType:
		startGame(ctx, req.GetMatchResult())
	case gameResultType:
		gameResult(req.GetGameResult())
	}
	return &resultPb.ResultResp{Message: "success"}, nil
}

// 调用snakeClient的start开始进行游戏
func startGame(ctx context.Context, matchResult *resultPb.MatchResult) {
	game, err := snake.StartGame(ctx, &snakePb.StartGameReq{
		AId:    matchResult.GetAId(),
		ABotId: matchResult.GetABotId(),
		BId:    matchResult.GetBId(),
		BBotId: matchResult.GetBBotId(),
	})
	if err != nil {
		zap.L().Error("startGame err:", zap.Error(err))
		return
	}
	zap.L().Debug(game.Message)
}

func gameResult(result *resultPb.GameResult) {
	// 到时候要传入公共空间，现在先暂时不做处理
	logger.SugarLogger.Debug(result)
}
