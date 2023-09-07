package model

import (
	"ChattyDiaryBot/internal/config"
	"context"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client
var redisCtx context.Context

func InitRedisDB() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.Config.Database.Redis.RedisAddr,
		Password: config.Config.Database.Redis.RedisPassword,
		DB:       config.Config.Database.Redis.RedisDB,
	})
	redisCtx = context.Background()
}
