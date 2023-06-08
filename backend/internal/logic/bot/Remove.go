package bot

import (
	"backend/conf/mysql"
	"backend/conf/redis"
	botPublic "backend/internal/handlers"
	"strconv"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

func RemoveService(botId, userId int) (*botPublic.RemoveBotResp, error) {
	var resp botPublic.RemoveBotResp
	var err error
	var ctx context.Context

	Bot := mysql.Q.Bot
	botSile, err := Bot.WithContext(ctx).Where(Bot.ID.Eq(int32(botId))).Limit(1).Find()
	if len(botSile) == 0 {
		err = errors.New("bot不存在或已被删除")
		return nil, err
	}

	bot := botSile[0]
	if int(bot.UserID) != userId {
		err = errors.New("您不是他的主人，没有权限删除这个Bot")
		return nil, err
	}

	// 删除数据库中和的和redis之中的
	_, err = Bot.WithContext(ctx).Where(Bot.ID.Eq(int32(botId))).Delete()
	if err != nil {
		return nil, err
	}

	go func() {
		redis.RDB.Del(ctx, "kob:bot:"+strconv.Itoa(userId))
	}()

	resp.Message = "success"
	return &resp, nil
}
