package store

import (
	"github.com/go-redis/redis"
	"school-helper/config"
)

var RedisClient *redis.Client

func init() {
	redisConfig, err := config.Configure.Map("redis")
	if err != nil {
		panic(err)
	}

	option := redis.Options{
		Addr:     redisConfig["addr"].(string),
		Password: redisConfig["password"].(string),
		DB:       redisConfig["db"].(int),
		PoolSize: 10,
	}
	RedisClient = redis.NewClient(&option)
}
