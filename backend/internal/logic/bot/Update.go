package bot

import (
	"backend/conf/mysql"
	"backend/conf/redis"
	botPublic "backend/internal/handlers"
	"backend/internal/models"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

func UpdateBotService(botId, userId int, title, code, description string) (*botPublic.UpdateBotResp, error) {
	// TODO: 对bot的信息做修改
	var err error
	var ctx context.Context

	Bot := mysql.Q.Bot

	forBotId := Bot.WithContext(ctx).Where(Bot.ID.Eq(int32(botId)))
	// 找到这个bot
	bot, err := forBotId.Limit(1).First()
	if errors.As(err, &gorm.ErrRecordNotFound) { // 表示没有找到
		err = errors.New("该bot不存在，可能已被删除")
		return nil, err
	}

	// 先判断这个bot是不是属于这个用户
	if int(bot.UserID) != userId {
		err = errors.New("您不是这个bot的主人，您没有权限修改它")
		return nil, err
	}

	// 判断title，code，description为不为空
	// TrimSpace 删除字符串两边的空格，中间的无法删除
	if strings.EqualFold(strings.TrimSpace(title), "") || len(title) == 0 {
		err = errors.New("标题不能为空")
		return nil, err
	}

	if len(title) > 100 {
		err = errors.New("标题长度不能超过100")
		return nil, err
	}

	// Replace(s, old, now, n)把字符串s中的old字符替换成new字符，n个数，-1表示全部
	if strings.EqualFold(strings.Replace(code, " ", "", -1), "") || len(code) == 0 {
		err = errors.New("代码不能为空")
		return nil, err
	}

	if len(code) > 10000 {
		err = errors.New("代码长度不能超过10000")
		return nil, err
	}

	if strings.EqualFold(strings.TrimSpace(description), "") || len(description) == 0 {
		description = "这个作者很懒。什么都没有写"
	}

	if len(description) > 300 {
		err = errors.New("Bot简介的长度不能超过300")
	}

	botRecord, _ := Bot.WithContext(ctx).Where(Bot.UserID.Eq(int32(userId))).Count()
	if botRecord > 10 {
		err = errors.New("bot的数量不能多于10个")
		return nil, err
	}

	botUpdate := models.Bot{
		ID:          bot.ID,
		UserID:      bot.UserID,
		Title:       title,
		Description: description,
		Code:        code,
		Createtime:  bot.Createtime,
		Modifytime:  time.Now(),
	}

	// 修改bot的信息，删除redis中这个用户的bot的信息  key:  kob:bot:`userId`

	go func() {
		redis.RDB.Del(ctx, "kob:bot:"+strconv.Itoa(userId))
	}()

	updateInfo, err := forBotId.Updates(botUpdate)
	log.Println(updateInfo)
	if err != nil {
		return nil, err
	}

	var resp botPublic.UpdateBotResp
	resp.Message = "success"
	return &resp, nil
}
