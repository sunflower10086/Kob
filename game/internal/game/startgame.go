package game

import (
	"context"
	"snake/conf/mysql"
	"snake/internal/game/util"
	resultPb "snake/internal/grpc/client/pb"
	"snake/internal/models"
	snakePb "snake/internal/pb"
	"snake/pkg/mw"
)

func StartGame(ctx context.Context, aId, aBotId, bId, bBotID int32) (*snakePb.StartGameResp, error) {
	mw.SugarLogger.Debug("StartGame function used")

	var User, Bot = mysql.Q.User, mysql.Q.Bot
	a, _ := User.WithContext(ctx).Where(User.ID.Eq(aId)).First()
	b, _ := User.WithContext(ctx).Where(User.ID.Eq(bId)).First()

	botA, botB := &models.Bot{}, &models.Bot{}

	if aBotId != -1 {
		botA, _ = Bot.WithContext(ctx).Where(Bot.ID.Eq(aBotId)).First()
	}
	if bBotID != -1 {
		botB, _ = Bot.WithContext(ctx).Where(Bot.ID.Eq(bBotID)).First()
	}

	Gamemap = NewGameMap(13, 14, 20, int(a.ID), botA, int(b.ID), botB)

	Gamemap.CreateMap()

	go Gamemap.Start()

	// 要严格对照上面的注释生成data
	data := make(map[string]interface{})
	data["AId"] = a.ID
	data["ASx"] = int32(13 - 2)
	data["ASy"] = int32(1)
	data["BId"] = b.ID
	data["BSx"] = int32(1)
	data["BSy"] = int32(14 - 2)

	var dataMap [13]resultPb.Edge
	for i, rows := range Gamemap.GetGameMap() {
		dataMap[i].Edge = rows
	}
	data["GameMap"] = &dataMap

	var playerA util.Player
	playerA.Id = int(a.ID)
	playerA.Username = a.Username
	playerA.Photo = a.Photo
	data["PlayerA"] = &playerA

	var playerB util.Player
	playerB.Id = int(b.ID)
	playerB.Username = b.Username
	playerB.Photo = b.Photo
	data["PlayerB"] = &playerB

	SendGameMap(data)

	var resp snakePb.StartGameResp
	resp.Message = "start game success"

	return &resp, nil
}
