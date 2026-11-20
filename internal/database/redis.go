package database

import (
	"context"

	"github.com/redis/go-redis/v9"

	"golang-restapi/internal/config"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func ConnectRedis() {

	RedisClient = redis.NewClient(
		&redis.Options{
			Addr:     config.GetEnv("REDIS_ADDR"),
			Password: config.GetEnv("REDIS_PASSWORD"),
			DB:       0,
		},
	)

	RedisClient.Ping(Ctx)
}
