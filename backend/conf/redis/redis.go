package redis

import (
	"backend/conf/settings"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func Init(conf *settings.AppConfig) error {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			conf.RedisConf.Host,
			conf.RedisConf.Port,
		),
		Password: conf.RedisConf.Password,
		DB:       conf.RedisConf.Db,
		PoolFIFO: false,
		// PoolTimeout 代表如果连接池所有连接都在使用中，等待获取连接时间，超时将返回错误
		// 默认是 1秒+ReadTimeout
		PoolTimeout: time.Duration(1 + 13),
	})

	RDB = rdb
	return nil
}
