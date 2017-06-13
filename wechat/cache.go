package wechat

import (
	"github.com/go-redis/redis"
	"time"
)

type Redis struct {
	conn *redis.Client
}

func NewRedis() *redis.Client {
	op := redis.Options{
		Addr:     "127.0.0.1:6389",
		Password: "luowei2008",
		DB:       0,
	}

	return redis.NewClient(&op)

}

//Get the value
func (redis *Redis) Get(key string) interface{} {

	if val, err := redis.conn.Get(key).Result(); err == nil {
		return val
	}
	return nil

}

//Set the key & value
func (redis *Redis) Set(key string, val interface{}, timeout time.Duration) error {
	err := redis.conn.Set(key, val, timeout).Err()
	if err != nil {
		return err
	}

	return nil

}

func (redis *Redis) Delete(key string) error {
	if err := redis.conn.Del(key).Err(); err != nil {
		return err
	}

	return nil

}

func (redis *Redis) IsExist(key string) bool {

	val, err := redis.conn.Exists(key).Result()
	if err != nil {
		return false
	}
	if val > 0 {
		return true
	} else {
		return false
	}

}
