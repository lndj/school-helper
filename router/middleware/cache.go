package middleware

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/go-redis/redis"
)

func InitRedis() gin.HandlerFunc {
	return func(c *gin.Context) {

		option := redis.Options{
			Addr:     "127.0.0.1:6389",
			Password: "luoning",
			DB:       0,
			PoolSize: 10,
		}
		//TODO This is useless！ Fuck！ just a string var
		redisClient := redis.NewClient(&option)

		// Set the Redis client to the context
		c.Set("redisClient", redisClient)

		// before request

		c.Next()

	}
}
