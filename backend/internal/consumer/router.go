package consumer

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
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
		log.Println(err)
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
	//log.Println("star-matching")
	//intBotId, _ := strconv.Atoi(botId)
	//intUserId, _ := strconv.Atoi(Clt.UserId)
	//
	//req := &matching_system.User{
	//	UserId: int32(intUserId),
	//	BotId:  int32(intBotId),
	//}
	//
	//_, err := match.AddUser(ctx, req)
	//if err != nil {
	//	handlers.SendResponse(ctx, err, nil)
	//	return
	//}
}

func stopMatching(ctx *gin.Context) {
	// TODO: 通过rpc去访问matchingSystem
	//log.Println("stop-matching")
	//
	//intUserId, _ := strconv.Atoi(Clt.UserId)
	//
	//req := &matching_system.User{
	//	UserId: int32(intUserId),
	//}
	//
	//_, err := match.RemoveUser(ctx, req)
	//if err != nil {
	//	handlers.SendResponse(ctx, err, nil)
	//}
}

func move(client *Client, d int) {
	// TODO: 设置下一步
	////klog.Infof("direction: ", d)
	//
	////TODO: 把移动的方向发送给game_system
	//game.Space.ClientDirection <- transform.Pair{
	//	PlayerId:  client.UserId,
	//	Direction: strconv.Itoa(d),
	//}
}
