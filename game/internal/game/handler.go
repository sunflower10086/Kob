package game

import (
	"encoding/json"
	"snake/internal/grpc/client"
	resultPb "snake/internal/grpc/client/pb"
	snakePb "snake/internal/pb"
	"snake/pkg/mw"
	"strconv"
	"sync"

	"go.uber.org/zap"
	"golang.org/x/net/context"
)

type SnakeImpl struct {
	snakePb.UnimplementedGameSystemServer
}

var (
	wg          sync.WaitGroup
	MoveMessage = make(chan *snakePb.SetNextStepResp, 100)
	Gamemap     *GameMap
)

func (SnakeImpl) StartGame(ctx context.Context, req *snakePb.StartGameReq) (res *snakePb.StartGameResp, err error) {

	// TODO: 业务逻辑
	mw.SugarLogger.Debugf("start game req: %v", req)
	return StartGame(ctx, req.GetAId(), req.GetABotId(), req.GetBId(), req.GetBBotId())
}

func (SnakeImpl) SetNextStep(ctx context.Context, req *snakePb.SetNextStepReq) (*snakePb.SetNextStepResp, error) {
	// TODO: 业务逻辑
	zap.L().Debug("SetNextStep function used")

	// 接收消息
	direction, _ := strconv.Atoi(req.GetDirection())
	mw.SugarLogger.Debug(req)
	Move(req.GetPlayerId(), int32(direction))

	resp := make(chan *snakePb.SetNextStepResp)
	go func() {
		for {
			select {
			case resp <- <-MoveMessage:
				return
			}
		}
	}()

	mw.SugarLogger.Debug(resp)

	return <-resp, nil
}

func Move(playerId string, direction int32) {
	PlayerId, _ := strconv.Atoi(playerId)

	switch PlayerId {
	case Gamemap.GetPlayerA().Id:
		if Gamemap.GetPlayerA().BotId == -1 { // 亲自出马
			Gamemap.SetNestStepA(direction)
		}
	case Gamemap.GetPlayerB().Id:
		if Gamemap.GetPlayerB().BotId == -1 { // 亲自出马
			Gamemap.SetNestStepB(direction)
		}
	}
}

func SendGameMap(data map[string]interface{}) {

	var gameMap resultPb.GameMap

	marshal, _ := json.Marshal(data)
	err := json.Unmarshal(marshal, &gameMap)
	if err != nil {
		mw.SugarLogger.Debugf("转化错误：", err)
		return
	}
	//klog.Infof("发送给前端的地图: ", &gameMap)

	resp := resultPb.ResultReq{
		EventType: 2,
		GameMap:   &gameMap,
	}

	_, err = client.Result(context.Background(), &resp)
	if err != nil {
		mw.SugarLogger.Debug(err)
	}
}
