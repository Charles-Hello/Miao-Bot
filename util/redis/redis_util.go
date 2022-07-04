package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"wechat_http/config"
)

var Rdb *redis.Client
var RdbCtx = context.Background()

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Host + ":" + config.Config.Redis.Port,
		Password: config.Config.Redis.Pwd,
		DB:       config.Config.Redis.Db,
	})
	Rdb = rdb

	err := rdb.Set(RdbCtx, "hello wechat", "redis初始化成功", 0).Err()
	if err != nil {
		panic(err)
	}
}
