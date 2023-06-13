package snake

import (
	"backend/conf/logger"
	pb "backend/internal/grpc/client/snake/pb"
	shape "backend/pkg/share_space"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

var (
	eg errgroup.Group
)

// WithSnake 只有一种消息，设置下一步，以便之后进行拓展
type WithSnake struct {
	Msg        chan shape.Pair
	GameClient pb.GameSystemClient
}

func (s *WithSnake) Send() error {
	streamCli, err := s.GameClient.SetNextStep(context.Background())
	if err != nil {
		return err
	}

	go ReadWebMsg()

	eg.Go(func() error {
		for {
			select {
			case msg := <-s.Msg:
				req := &pb.SetNextStepReq{Direction: msg.Direction, PlayerId: msg.PlayerId}
				logger.SugarLogger.Debugf("move: %v", req)

				if err := streamCli.Send(req); err != nil {
					return err
				}
			}
			time.Sleep(time.Second)
		}
	})

	return eg.Wait()
}

func (s *WithSnake) Receive() error {
	streamCli, err := s.GameClient.SetNextStep(context.Background())
	if err != nil {
		return err
	}

	eg.Go(func() error {
		for {
			zap.L().Debug("Receive")
			if resp, err := streamCli.Recv(); err != nil {
				if err == io.EOF {
					zap.L().Debug("receive done")
				}
			} else {
				logger.SugarLogger.Debugf("Receive Message " + resp.Event)

				if strings.EqualFold("发送地图", resp.Event) {
					// 在公共空间中传入信息
					GetMap(resp)
				}

				if strings.EqualFold("move", resp.Event) {
					getMove(resp)
				}
				//// 我们每次从后端获得消息都要关闭一下，防止snake服务断开之后我们这个也不能用
				//streamCli.CloseSend()
			}

			time.Sleep(time.Second)
		}
	})

	return eg.Wait()
}

func GetMap(resp *pb.SetNextStepResp) {
	// TODO: 必须把resp转换为SnakeGame
	gameMap := resp.GetGameMap()
	sendGame(gameMap)
}

// 获得后端传来的地图，传给前端
func sendGame(gameMap *pb.GameMap) {
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

	Space.Game <- respMap
}

// ReadWebMsg 死循环读取前端传来的消息
func ReadWebMsg() {
	for {
		select {
		case msg := <-Space.ClientDirection:
			SnakeMd.Msg <- msg
		}
	}
}

func getMove(resp *pb.SetNextStepResp) {
	logger.SugarLogger.Debug(resp)

	fmt.Println(resp.GetADirection(), resp.GetBDirection())

	pair := shape.Pair{
		Event:      "move",
		ADirection: strconv.Itoa(int(resp.GetADirection())),
		BDirection: strconv.Itoa(int(resp.GetBDirection())),
	}
	logger.SugarLogger.Debug(pair)

	Space.ServiceDirection <- pair
}
