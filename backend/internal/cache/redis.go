package cache

import (
	"context"
	"math"
	"os"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	once        sync.Once
)

func GetRedisClient() *redis.Client {
	once.Do(func() {
		addr := os.Getenv("REDIS_ADDR")
		if addr == "" {
			addr = "localhost:6379"
		}
		redisClient = redis.NewClient(&redis.Options{
			Addr: addr,
		})
	})
	return redisClient
}

func SetSlug(ctx context.Context, slug, url string, expiresAt *time.Time) error {
	client := GetRedisClient()
	
	d := time.Until(*expiresAt)
	ttl := time.Duration(math.Max(float64(d), 0))
	if ttl > 0 {
		return client.Set(ctx, slug, url, ttl).Err()
	}
	return nil
}

func GetSlug(ctx context.Context, slug string) (string, error) {
	client := GetRedisClient()
	return client.Get(ctx, slug).Result()
}
