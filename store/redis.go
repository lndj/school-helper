package store

import (
	"github.com/go-redis/redis"
	"school-helper/config"
)

var RedisClient *redis.Client

func InitRedisClient() error {

	redisConfig, err := config.Configure.Map("redis")
	if err != nil {
		return err
	}

	option := redis.Options{
		Addr:     string(redisConfig["addr"]),
		Password: string(redisConfig["password"]),
		DB:       int(redisConfig["db"]),
		PoolSize: 10,
	}
	RedisClient = redis.NewClient(&option)

	return nil
}
