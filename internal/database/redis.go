package database

import (
	"context"
	"fmt"

	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func NewRedisClient() (*redis.Client, error) {

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	// ทดสอบการเชื่อมต่อ (Ping)
	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect redis: %v", err)
	}

	return rdb, nil
}
