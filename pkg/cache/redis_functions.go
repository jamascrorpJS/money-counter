package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	redis *redis.Client
}

func NewRedisClient() *Redis {
	return &Redis{
		redis: RedisClient(),
	}
}

func (r *Redis) SetKey(ctx context.Context, key string, val interface{}, ttl time.Duration) error {
	return r.redis.Set(ctx, key, val, ttl).Err()
}

func (r *Redis) GetValue(ctx context.Context, key string) (string, error) {
	return r.redis.Get(ctx, key).Result()
}

func (r *Redis) Del(ctx context.Context, key string) {
	r.redis.Del(ctx, key)
}
