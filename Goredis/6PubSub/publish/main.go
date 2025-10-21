package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "845924",
		DB:       0,
	})
	ctx := context.Background()

	// 发布
	rdb.Publish(ctx, "channel1", "msg")
}
