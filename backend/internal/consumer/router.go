package consumer

import (
	"backend/internal/grpc/client/match"
	pb "backend/internal/grpc/client/match/pb"
	"backend/internal/grpc/client/snake"
	"backend/pkg/result"
	shape "backend/pkg/share_space"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"go.uber.org/zap"
)

type Req struct {
	Event     string `json:"event"`
	BotId     string `json:"bot_id,omitempty"`
	Direction int    `json:"direction,omitempty"`
}

var Clt *Client

func Router(ctx *gin.Context, client *Client, message string) {
	Clt = client
	var data Req
	err := json.Unmarshal([]byte(message), &data)
	if err != nil {
		zap.L().Error(err.Error())
		return
	}

	event := data.Event

	if strings.EqualFold(event, "start-matching") {
		starMatching(ctx, data.BotId)
	} else if strings.EqualFold(event, "stop-matching") {
		stopMatching(ctx)
	} else if strings.EqualFold(event, "move") {
		move(client, data.Direction)
	}
}

func starMatching(ctx *gin.Context, botId string) {
	// TODO: 通过rpc去访问matchingSystem
	zap.L().Debug("star-matching")
	intBotId, _ := strconv.Atoi(botId)
	intUserId, _ := strconv.Atoi(Clt.UserId)

	req := &pb.User{
		UserId: int32(intUserId),
		BotId:  int32(intBotId),
	}

	_, err := match.AddUser(ctx, req)
	if err != nil {
		result.SendResult(ctx, result.Fail(err))
		return
	}
}

func stopMatching(ctx *gin.Context) {
	// TODO: 通过rpc去访问matchingSystem
	zap.L().Debug("stop-matching")

	intUserId, _ := strconv.Atoi(Clt.UserId)

	req := &pb.User{
		UserId: int32(intUserId),
	}

	_, err := match.RemoveUser(ctx, req)
	if err != nil {
		result.SendResult(ctx, result.Fail(err))
	}
}

// 获得前端传来的移动信息(下一步的移动信息，把信息发送到公共区域)
func move(client *Client, d int) {
	////TODO: 把移动的方向发送给game_system
	snake.Space.ClientDirection <- shape.Pair{
		PlayerId:  client.UserId,
		Direction: strconv.Itoa(d),
	}
}
