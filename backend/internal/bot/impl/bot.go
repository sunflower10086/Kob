package impl

import (
	"backend/conf/mysql"
	"backend/conf/redis"
	"backend/internal"
	"backend/internal/bot"
	"backend/internal/models"
	"backend/pkg/ioc"
	"encoding/json"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	impl = &BotServiceImpl{}
)

func init() {
	ioc.RegistryImpl(impl)
}

type BotServiceImpl struct {
	l *log.Logger
}

func (b *BotServiceImpl) Name() string {
	return internal.BotServiceName
}

func (b *BotServiceImpl) Config() {
	b.l = log.New(os.Stderr, "  [bot] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func NewHostServiceImpl() *BotServiceImpl {
	return &BotServiceImpl{
		l: log.New(os.Stderr, "  [bot] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (b *BotServiceImpl) AddBot(ctx *gin.Context, botReq *bot.AddBotReq) (*bot.AddBotResp, error) {
	var resp bot.AddBotResp
	var err error

	title, code, description, userId := botReq.Title, botReq.Code, botReq.Description, botReq.UserId

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
		err = errors.New("bot简介的长度不能超过300")
	}

	Bot := mysql.Q.Bot
	botRecord, _ := Bot.WithContext(ctx).Where(Bot.UserID.Eq(int32(userId))).Count()
	if botRecord > 10 {
		err = errors.New("bot的数量不能多于10个")
		return nil, err
	}

	sqlBot := models.Bot{
		UserID:      int32(userId),
		Title:       title,
		Description: description,
		Code:        code,
		Createtime:  time.Now(),
		Modifytime:  time.Now(),
	}
	err = Bot.WithContext(ctx).Create(&sqlBot)
	if err != nil {
		return nil, err
	}

	resp.Message = "success"
	return &resp, nil
}

func (b *BotServiceImpl) GetListBot(ctx *gin.Context, req *bot.GetListBotReq) (*bot.GetListBotResp, error) {
	var resp bot.GetListBotResp

	//先去redis中尝试获取 key: kob:bot:`userId`
	redisBotList, err := redis.RDB.Get(ctx, "cache:kob:bot:"+req.UserId).Result()
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

	userId, _ := strconv.Atoi(req.UserId)
	//直接获取
	Bot := mysql.Q.Bot
	botList, err := Bot.WithContext(ctx).Where(Bot.UserID.Eq(int32(userId))).Find()
	if err != nil {
		return nil, err
	}

	for _, sqlBot := range botList {
		temp := bot.ResultBot{
			ID:          sqlBot.ID,
			UserID:      sqlBot.UserID,
			Title:       sqlBot.Title,
			Description: sqlBot.Description,
			Code:        sqlBot.Code,
			CreateTime:  sqlBot.Createtime.Format("2006-01-02 15:04:05"),
			ModifyTime:  sqlBot.Modifytime.Format("2006-01-02 15:04:05"),
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

func (b *BotServiceImpl) UpdateBot(ctx *gin.Context, req *bot.UpdateBotReq) (*bot.UpdateBotResp, error) {
	var err error
	strUserId, strBotId, title, code, description := req.UserId, req.BotId, req.Title, req.Code, req.Description

	botId, _ := strconv.Atoi(strBotId)
	userId, _ := strconv.Atoi(strUserId)

	Bot := mysql.Q.Bot

	forBotId := Bot.WithContext(ctx).Where(Bot.ID.Eq(int32(botId)))
	// 找到这个bot
	sqlBot, err := forBotId.Limit(1).First()
	if errors.As(err, &gorm.ErrRecordNotFound) { // 表示没有找到
		err = errors.New("该bot不存在，可能已被删除")
		return nil, err
	}

	// 先判断这个bot是不是属于这个用户
	if int(sqlBot.UserID) != userId {
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
		err = errors.New("bot简介的长度不能超过300")
	}

	botRecord, _ := Bot.WithContext(ctx).Where(Bot.UserID.Eq(int32(userId))).Count()
	if botRecord > 10 {
		err = errors.New("bot的数量不能多于10个")
		return nil, err
	}

	botUpdate := models.Bot{
		ID:          sqlBot.ID,
		UserID:      sqlBot.UserID,
		Title:       title,
		Description: description,
		Code:        code,
		Createtime:  sqlBot.Createtime,
		Modifytime:  time.Now(),
	}

	// 修改bot的信息，删除redis中这个用户的bot的信息  key:  kob:bot:`userId`

	go func() {
		redis.RDB.Del(ctx, "kob:bot:"+strconv.Itoa(userId))
	}()

	_, err = forBotId.Updates(botUpdate)
	if err != nil {
		return nil, err
	}

	var resp bot.UpdateBotResp
	resp.Message = "success"
	return &resp, nil
}

func (b *BotServiceImpl) DeleteBot(ctx *gin.Context, req *bot.DeleteBotReq) (*bot.DeleteBotResp, error) {
	var resp bot.DeleteBotResp
	var err error

	botId, _ := strconv.Atoi(req.BotId)
	userId := req.UserId

	Bot := mysql.Q.Bot
	botSile, err := Bot.WithContext(ctx).Where(Bot.ID.Eq(int32(botId))).Limit(1).Find()
	if len(botSile) == 0 {
		err = errors.New("bot不存在或已被删除")
		return nil, err
	}

	sqlBot := botSile[0]
	if int(sqlBot.UserID) != userId {
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
