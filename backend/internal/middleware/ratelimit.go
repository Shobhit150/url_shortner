package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/ulule/limiter/v3"
	ginmiddleware "github.com/ulule/limiter/v3/drivers/middleware/gin"
	redisstore "github.com/ulule/limiter/v3/drivers/store/redis"
)


func RateLimiter() gin.HandlerFunc {
	
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "",
        DB:       0,
	})
	rate, _ := limiter.NewRateFromFormatted("10-M")

	store, err := redisstore.NewStoreWithOptions(rdb, limiter.StoreOptions{
		Prefix: "rate_limiter",
	})
	if err != nil {
		panic(err)
	}

	lim := limiter.New(store, rate)

	return ginmiddleware.NewMiddleware(lim)
}