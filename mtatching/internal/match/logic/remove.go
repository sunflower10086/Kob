package logic

import (
	"matching/internal/match/logic/matchutil"
	pb "matching/internal/pb/matchingServer"
	"matching/pkg/mw"
	"strconv"

	"go.uber.org/zap"
	"golang.org/x/net/context"
)

func Remove(ctx context.Context, userId int32) (*pb.Response, error) {
	zap.L().Debug("Remove func used")
	// 其他线程调用这个函数的时候我们这个线程本身可能也会调用这个players可能读写冲突所以要加锁
	lock.Lock()
	defer lock.Unlock()

	newPlayer := make([]matchutil.Player, 0)
	for i := 0; i < len(matchutil.Players); i++ {
		if matchutil.Players[i].UserId != userId {
			newPlayer = append(newPlayer, matchutil.Players[i])
		}
	}

	matchutil.Players = newPlayer

	mw.SugarLogger.Debugf("Players: %v", matchutil.Players)

	var resp pb.Response
	resp.Message = "remove user" + strconv.Itoa(int(userId))

	zap.L().Debug(resp.Message)

	return &resp, nil
}
