package bot

import (
	"backend/conf/mysql"
	"backend/conf/redis"
	botPublic "backend/internal/handlers"
	"strconv"
	"time"

	"github.com/goccy/go-json"
	"golang.org/x/net/context"
)

func GetListService(userId int) (*botPublic.GetListBotResp, error) {
	// TODO: 获得用户的bot列表
	var resp botPublic.GetListBotResp
	ctx := context.Background()

	//先去redis中尝试获取 key: kob:bot:`userId`
	redisBotList, err := redis.RDB.Get(ctx, "cache:kob:bot:"+strconv.Itoa(userId)).Result()
	if err != nil {
		if len(redisBotList) != 0 {
			err := json.Unmarshal([]byte(redisBotList), &resp.BotList)
			if err != nil {
				return nil, err
			}
			return &resp, nil
		}
	}
	// 获取到直接返回，没有得到再去数据库中获取

	//直接获取
	Bot := mysql.Q.Bot
	botList, err := Bot.WithContext(ctx).Where(Bot.UserID.Eq(int32(userId))).Find()
	if err != nil {
		return nil, err
	}

	for _, bot := range botList {
		temp := botPublic.ResultBot{
			ID:          bot.ID,
			UserID:      bot.UserID,
			Title:       bot.Title,
			Description: bot.Description,
			Code:        bot.Code,
			CreateTime:  bot.Createtime.Format("2006-01-02 15:04:05"),
			ModifyTime:  bot.Modifytime.Format("2006-01-02 15:04:05"),
		}
		resp.BotList = append(resp.BotList, &temp)
	}

	// 往redis中添加这个缓存  key: kob:bot:`userId`，在redis中的类型是
	if len(resp.BotList) > 0 {
		go func() {
			for i := 0; i < 5; i++ {
				marshal, err := json.Marshal(resp.BotList)
				if err != nil {
					continue
				}
				redis.RDB.Set(ctx, "cache:kob:bot:"+strconv.Itoa(userId), marshal, time.Minute*30)
			}
		}()
	}

	return &resp, nil
}
