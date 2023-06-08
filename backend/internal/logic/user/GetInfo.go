package user

import (
	"backend/conf/logger"
	"backend/conf/mysql"
	"backend/conf/redis"
	userPublic "backend/internal/handlers"
	"backend/internal/models"
	"context"
	"strconv"
	"time"

	"github.com/goccy/go-json"
)

func GetInfoService(userId int) (*userPublic.GetUserInfoResponse, error) {
	// TODO: 获取用户信息的代码

	var resp userPublic.GetUserInfoResponse
	ctx := context.Background()

	// 1.先去redis中查，查到直接返回
	userRedis, _ := redis.RDB.Get(ctx, "cache:kob:user:"+strconv.Itoa(userId)).Result()
	if len(userRedis) != 0 {
		var user models.User
		err := json.Unmarshal([]byte(userRedis), &user)
		if err != nil {
			return &resp, err
		}

		resp.UserId = string(user.ID)
		resp.Username = user.Username
		resp.Photo = user.Photo
		logger.SugarLogger.Debugf("GetInfoHandler in redis %v", resp)
		return &resp, nil
	}

	// 2.不存在先去数据库中查，查到之后写入redis
	var User = mysql.Q.User
	user, err := User.WithContext(ctx).Where(User.ID.Eq(int32(userId))).Limit(1).Find()
	if err != nil {
		return &resp, err
	}

	// 3.存入redis
	go func() {
		userJson, err := json.Marshal(user[0])
		if err != nil {
			return
		}
		redis.RDB.Set(ctx, "cache:kob:user:"+strconv.Itoa(userId), userJson, time.Second*60*60*24)
	}()

	resp.UserId = string(user[0].ID)
	resp.Username = user[0].Username
	resp.Photo = user[0].Photo

	return &resp, err
}
