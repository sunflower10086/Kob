package logic

import (
	"matching/conf/mysql"
	"matching/internal/match/logic/matchutil"
	pb "matching/internal/pb/matchingServer"
	"strconv"
	"sync"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

var (
	lock = sync.Mutex{}
)

func AddUser(ctx context.Context, userId, botId int32) (*pb.Response, error) {
	zap.L().Debug("AddUser func used")

	var User = mysql.Q.User
	user, err := User.WithContext(ctx).Where(User.ID.Eq(userId)).First()
	if errors.As(err, &gorm.ErrRecordNotFound) { // 没有找到
		return nil, errors.New("数据库中没有这个用户")
	}

	// 其他线程调用这个函数的时候我们这个线程本身可能也会调用这个players可能读写冲突所以要加锁
	lock.Lock()
	defer lock.Unlock()

	matchutil.Players = append(matchutil.Players, matchutil.Player{UserId: userId, BotId: botId, Rating: *user.Rating, WaitTime: 0})

	// 返回匹配成功的userId与他们的botId
	var resp pb.Response
	resp.Message = "add user success " + strconv.Itoa(int(userId)) + " " + strconv.Itoa(int(botId))

	zap.L().Debug(resp.Message)
	return &resp, nil
}
