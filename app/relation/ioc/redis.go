package ioc

import (
	"github.com/redis/go-redis/v9"

	"github.com/crazyfrankie/douyin/app/relation/config"
)

func InitRedis() redis.Cmdable {
	return redis.NewClient(&redis.Options{
		Addr: config.GetConf().Redis.Address,
	})
}
