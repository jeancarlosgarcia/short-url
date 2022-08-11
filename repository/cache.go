package repository

import (
	"context"
	"time"

	"short-url/interfaces"

	"github.com/go-redis/redis/v8"
)

type CacheRepository struct {
	redis *redis.Client
}

func NewCacheRepository(redis *redis.Client) interfaces.ICache {
	return &CacheRepository{ redis: redis }
}

func (c CacheRepository) Set(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	return c.redis.Set(ctx, key, value, exp).Err()
}

func (c CacheRepository) Get(ctx context.Context, key string) (string, error) {
	return c.redis.Get(ctx, key).Result()
}

func (c CacheRepository) Del(ctx context.Context, key string) error {
	return c.redis.Del(ctx, key).Err()
}