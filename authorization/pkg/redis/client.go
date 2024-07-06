package redis

import (
	"authorization/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(config config.Redis) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       config.DB,
	})
}
