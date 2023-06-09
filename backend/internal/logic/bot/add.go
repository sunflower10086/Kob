package bot

import (
	"backend/conf/mysql"
	botPublic "backend/internal/handlers"
	"backend/internal/models"
	"strings"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

func AddBotService(title, code, description string, userId int) (*botPublic.AddBotResp, error) {
	var resp botPublic.AddBotResp
	var err error
	ctx := context.Background()

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

	Bot := mysql.Q.Bot
	botRecord, _ := Bot.WithContext(ctx).Where(Bot.UserID.Eq(int32(userId))).Count()
	if botRecord > 10 {
		err = errors.New("bot的数量不能多于10个")
		return nil, err
	}

	bot := models.Bot{
		UserID:      int32(userId),
		Title:       title,
		Description: description,
		Code:        code,
		Createtime:  time.Now(),
		Modifytime:  time.Now(),
	}
	err = Bot.WithContext(ctx).Create(&bot)
	if err != nil {
		return nil, err
	}

	resp.Message = "success"
	return &resp, nil
}
