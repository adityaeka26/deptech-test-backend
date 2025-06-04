package redis

import (
	"context"

	"github.com/adityaeka26/deptech-test-backend/config"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	RedisClient *redis.Client
}

func NewRedis(config *config.EnvConfig) (*Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: config.RedisPassword,
		DB:       0,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &Redis{
		RedisClient: rdb,
	}, nil
}
