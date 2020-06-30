package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

func DistributedLock(ctx context.Context, client *redis.Client, key string, value interface{}, timeout time.Duration) (bool, error) {
	return client.SetNX(ctx, key, value, timeout).Result()
}

func DistributedUnlock(ctx context.Context, client *redis.Client, key string) (int64, error) {
	return client.Del(ctx, key).Result()
}

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
