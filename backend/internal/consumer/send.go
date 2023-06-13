package consumer

import (
	"backend/conf/logger"
	"backend/internal/grpc/client/snake"
	shape "backend/pkg/share_space"
	"encoding/json"

	"go.uber.org/zap"
)

type SnakeGame struct {
	AId string    `json:"a_id"`
	ASx string    `json:"a_sx"`
	ASy string    `json:"a_sy"`
	BId string    `json:"b_id"`
	BSx string    `json:"b_sx"`
	BSy string    `json:"b_sy"`
	Map [][]int32 `json:"map"`
}

// SendMsg 从通道公共空间中取出来
func (c *Client) SendMsg() {
	for {
		select {
		case data := <-snake.Space.Game:
			sendMap(data)
		case data := <-snake.Space.ServiceDirection:
			logger.SugarLogger.Debug(data)
			sendDirection(data)
		case data := <-snake.Space.Result:
			sendResult(data)
		}
	}
}

func sendMap(data *shape.SnakeGame) {

	AId, BId := data.AId, data.BId

	respGame := SnakeGame{
		AId: data.AId,
		ASx: data.ASy,
		ASy: data.ASx,
		BId: data.BId,
		BSx: data.BSx,
		BSy: data.BSy,
		Map: data.Map,
	}

	// 给用户发送消息
	respA, err := marshal(NewRespForWeb(
		WithEvent("start-game"),
		WithOpponentUsername(data.PlayerB.Username),
		WithOpponentPhoto(data.PlayerB.Photo),
		WithGame(respGame),
	))
	if err != nil {
		zap.L().Debug("marshal: ", zap.Error(err))
		return
	}

	respB, err := marshal(NewRespForWeb(
		WithEvent("start-game"),
		WithOpponentUsername(data.PlayerA.Username),
		WithOpponentPhoto(data.PlayerA.Photo),
		WithGame(respGame),
	))
	if err != nil {
		zap.L().Debug("marshal: ", zap.Error(err))
		return
	}

	for _, user := range Clt.Hub.Clients {
		// 给玩家A发送
		if user.UserId == AId {
			user.Send <- respA
		}

		// 给玩家B发送
		if user.UserId == BId {
			user.Send <- respB
		}
	}
}

func marshal(web *RespForWeb) ([]byte, error) {
	marshal, err := json.Marshal(web)
	if err != nil {
		return []byte{}, err
	}

	return marshal, nil
}

func sendDirection(data shape.Pair) {
	marshal, err := json.Marshal(data)
	if err != nil {
		zap.L().Debug("marshal: ", zap.Error(err))
		return
	}

	Clt.Hub.broadcast <- marshal
}

func sendResult(data shape.Result) {
	marshal, err := json.Marshal(data)
	if err != nil {
		zap.L().Debug("marshal: ", zap.Error(err))
		return
	}

	Clt.Hub.broadcast <- marshal
}
