package game

import (
	"encoding/json"
	"io"
	pb "snake/internal/pb"
	"snake/pkg/mw"
	"strconv"
	"time"

	"go.uber.org/zap"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

type SnakeImpl struct {
	pb.UnimplementedGameSystemServer
}

var (
	eg      errgroup.Group
	Message = make(chan *pb.SetNextStepResp, 100)
	Gamemap *GameMap
)

func (SnakeImpl) StartGame(ctx context.Context, req *pb.StartGameReq) (res *pb.StartGameResp, err error) {

	// TODO: 业务逻辑
	mw.SugarLogger.Debugf("start game req: %v", req)
	return StartGame(ctx, req.GetAId(), req.GetABotId(), req.GetBId(), req.GetBBotId())
}

func (SnakeImpl) SetNextStep(stream pb.GameSystem_SetNextStepServer) (err error) {
	// TODO: 业务逻辑
	zap.L().Debug("SetNextStep function used")

	// 接收消息
	eg.Go(func() error {
		for {
			req, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return err
			}
			mw.SugarLogger.Debugf("message received: %v\n", req)
			direction, _ := strconv.Atoi(req.Direction)

			Move(req.PlayerId, int32(direction))
			time.Sleep(time.Second)
		}
	})

	// 发送消息
	eg.Go(func() error {
		for {
			select {
			case msg := <-Message:
				mw.SugarLogger.Debugf("msg Event: %s", msg.GetEvent())
				err := stream.Send(msg)
				if err != nil {
					mw.SugarLogger.Debugf("err: %v", err)
					return err
				}
			}
		}
	})

	return eg.Wait()
}

func Move(playerId string, direction int32) {
	PlayerId, _ := strconv.Atoi(playerId)

	mw.SugarLogger.Debug(Gamemap.GetPlayerA())
	mw.SugarLogger.Debug(Gamemap.GetPlayerB())
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

	var gameMap pb.GameMap

	marshal, _ := json.Marshal(data)
	err := json.Unmarshal(marshal, &gameMap)
	if err != nil {
		mw.SugarLogger.Debugf("转化错误：", err)
		return
	}
	//klog.Infof("发送给前端的地图: ", &gameMap)

	resp := pb.SetNextStepResp{
		Event:      "发送地图",
		ADirection: 0,
		BDirection: 0,
		GameMap:    &gameMap,
	}

	Message <- &resp
}
