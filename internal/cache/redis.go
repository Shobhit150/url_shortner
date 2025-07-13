package cache

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	once sync.Once
)

func GetRedisClient() *redis.Client {
	once.Do(func ()  {
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

func SetSlug(ctx context.Context, slug, url string) error {
	client := GetRedisClient()
	return client.Set(ctx,slug,url, 24*time.Hour).Err()
}

func GetSlug(ctx context.Context, slug string) (string, error) {
	client := GetRedisClient()
	return client.Get(ctx,slug).Result()
}