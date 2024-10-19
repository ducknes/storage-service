package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"storage-service/settings"
)

func NewRedisClient(ctx context.Context, redisSettings settings.RedisSettings) (*redis.Client, error) {
	options := &redis.Options{
		Addr:     redisSettings.Address,
		Password: redisSettings.Password,
		DB:       redisSettings.Database,
	}

	client := redis.NewClient(options)
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
