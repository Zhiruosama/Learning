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

	// HSET
	// err := rdb.HSet(ctx, "user_1", "username", "zhangsan").Err()
	// if err != nil {
	// 	panic(err)
	// }

	// HGET
	// username, _ := rdb.HGet(ctx, "user_1", "username").Result()
	// fmt.Println(username)

	// HIncryBy
	// _ = rdb.HSet(ctx, "user_1", "count", 0).Err()
	// count, _ := rdb.HIncrBy(ctx, "user_1", "count", 2).Result()
	// fmt.Println(count)

	// HKEYS
	// keys, _ := rdb.HKeys(ctx, "user_1").Result()
	// fmt.Println(keys)

	// HDEL
	rdb.HDel(ctx, "user_1")

}
