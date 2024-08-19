package Service

import (
	_ "embed"
	"github.com/redis/go-redis/v9"
)

////go:embed test
//var food string

var redisInstance *redis.Client

func getRedis() *redis.Client {
	if redisInstance == nil {
		redisInstance = redis.NewClient(&redis.Options{
			Addr:     "192.168.10.83:6379",
			Password: "",
			DB:       0,
		})
	}
	return redisInstance
}
