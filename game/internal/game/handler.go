package game

import (
	"encoding/json"
	"fmt"
	"snake/internal/grpc/client/result"
	resultPb "snake/internal/grpc/client/result/pb"
	snakePb "snake/internal/pb"
	"snake/pkg/mw"
	"strconv"

	"golang.org/x/net/context"
)

type SnakeImpl struct {
	snakePb.UnimplementedGameSystemServer
}

var (
	Gamemap *GameMap
)

func (SnakeImpl) StartGame(ctx context.Context, req *snakePb.StartGameReq) (res *snakePb.StartGameResp, err error) {

	// TODO: 业务逻辑
	mw.SugarLogger.Debugf("start game req: %v", req)
	return StartGame(ctx, req.GetAId(), req.GetABotId(), req.GetBId(), req.GetBBotId())
}

func (SnakeImpl) SetNextStep(ctx context.Context, req *snakePb.SetNextStepReq) (*snakePb.SetNextStepResp, error) {
	// TODO: 业务逻辑
	// 接收消息
	direction, _ := strconv.ParseInt(req.GetDirection(), 10, 32)
	if req.IsCode {
		fmt.Println(direction)
		PlayerId, _ := strconv.Atoi(req.PlayerId)
		switch PlayerId {
		case Gamemap.GetPlayerA().Id:
			Gamemap.SetNestStepA(int32(direction))
		case Gamemap.GetPlayerB().Id:
			Gamemap.SetNestStepB(int32(direction))
		}
	}

	Move(req.GetPlayerId(), int32(direction), req.IsCode)

	resp := make(chan *snakePb.SetNextStepResp, 2)
	go func() {
		for {
			select {
			case message := <-Gamemap.MoveMessage:
				fmt.Println(req.PlayerId, req.IsCode)
				fmt.Println("send message is", message)
				resp <- message
				return
			}
		}
	}()

	return <-resp, nil
}

func Move(playerId string, direction int32, isBot bool) {
	PlayerId, _ := strconv.Atoi(playerId)

	switch PlayerId {
	case Gamemap.GetPlayerA().Id:
		if Gamemap.GetPlayerA().BotId == -1 { // 亲自出马
			Gamemap.SetNestStepA(direction)
		}
		if isBot {
			Gamemap.SetNestStepA(direction)
		}
	case Gamemap.GetPlayerB().Id:
		if Gamemap.GetPlayerB().BotId == -1 { // 亲自出马
			Gamemap.SetNestStepB(direction)
		}
		if isBot {
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

	resp := resultPb.ResultReq{
		EventType: 2,
		GameMap:   &gameMap,
	}

	_, err = result.Result(context.Background(), &resp)
	if err != nil {
		mw.SugarLogger.Debug(err)
	}
}
