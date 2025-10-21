package main

import (
	"context"
	"fmt"
	"time"

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

	err := rdb.Set(ctx, "rediskey", "redisvalue", time.Second*10).Err()
	if err != nil {
		panic(err)
	}
	oldval, err := rdb.GetSet(ctx, "rediskey", "new value").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(oldval)

	// 其实就是跟redis基本操作几乎一摸一样
	// rdb.SetNX()
	// rdb.MGet()
	// rdb.MSet()
	// rdb.Incr() 自增
	err = rdb.Set(ctx, "testnum", 0, 0).Err()
	if err != nil {
		panic(err)
	}
	getval, _ := rdb.Get(ctx, "testnum").Result()
	fmt.Println(getval)
	incrval, _ := rdb.Incr(ctx, "testnum").Result()
	fmt.Println(incrval)
}
