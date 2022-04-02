package database

import (
	"time"

	"github.com/go-redis/redis"
)

// RedisDBEngine RedisDBEngine
type RedisDBEngine struct {
	DBConn *redis.Client
}

// Get key key typ 1 字符串  2 数组 3 map
func (redis *RedisDBEngine) Get(key string, typ int) (interface{}, error) {

	switch typ {
	case 1:
		return redis.DBConn.Get(key).Result()
	case 2:
		return redis.DBConn.HMGet(key).Result()

	default:
		return redis.DBConn.Get(key).Result()
	}
}

// Set Set
func (redis *RedisDBEngine) Set(key string, value ...interface{}) error {

	_, err := redis.DBConn.Set(key, value, time.Minute*24*60).Result()
	return err
}
