package store

import "github.com/go-redis/redis"

var RedisClient *redis.Client

func init() {
	option := redis.Options{
		Addr:     "127.0.0.1:6389",
		Password: "luoning",
		DB:       0,
		PoolSize: 10,
	}
	RedisClient = redis.NewClient(&option)
}
