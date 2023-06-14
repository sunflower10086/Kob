package result

import (
	"backend/conf/logger"
	"backend/internal/grpc/client/snake"
	snakePb "backend/internal/grpc/client/snake/pb"
	resultPb "backend/internal/grpc/server/result/pb"
	shape "backend/pkg/share_space"
	"strconv"

	"go.uber.org/zap"
	"golang.org/x/net/context"
)

const (
	startGameType int32 = iota
	gameResultType
	gameMapType
)

type ResultServerImpl struct {
	resultPb.UnimplementedResultServer
}

func (r *ResultServerImpl) Result(ctx context.Context, req *resultPb.ResultReq) (*resultPb.ResultResp, error) {
	switch req.GetEventType() {
	case startGameType:
		startGame(ctx, req.GetMatchResult())
	case gameResultType:
		gameResult(req.GetGameResult())
	case gameMapType:
		getMap(req.GetGameMap())
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
	resp := shape.Result{
		Event: "result",
		Loser: result.GetLoser(),
	}
	zap.L().Debug(resp.Loser)
	snake.Space.Result <- resp
}

func getMap(gameMap *resultPb.GameMap) {
	// TODO: 必须把resp转换为SnakeGame
	// 把后端传来的游戏信息做一下格式转换
	Map := make([][]int32, len(gameMap.GetGameMap()))
	for i, edge := range gameMap.GetGameMap() {
		item := make([]int32, len(edge.Edge))
		for j, point := range edge.Edge {
			item[j] = point
		}
		Map[i] = append(Map[i], item...)
	}

	playerA := shape.Player{
		Photo:    gameMap.GetPlayerA().GetPhoto(),
		Username: gameMap.GetPlayerA().GetUsername(),
		UserID:   gameMap.GetPlayerA().GetUserID(),
	}

	playerB := shape.Player{
		Photo:    gameMap.GetPlayerB().GetPhoto(),
		Username: gameMap.GetPlayerB().GetUsername(),
		UserID:   gameMap.GetPlayerB().GetUserID(),
	}

	respMap := shape.NewSnakeGame(
		strconv.Itoa(int(gameMap.GetAId())),
		strconv.Itoa(int(gameMap.GetASx())),
		strconv.Itoa(int(gameMap.GetASy())),
		strconv.Itoa(int(gameMap.GetBId())),
		strconv.Itoa(int(gameMap.GetBSx())),
		strconv.Itoa(int(gameMap.GetBSy())),
		Map,
		playerA,
		playerB,
	)

	snake.Space.Game <- respMap
}
