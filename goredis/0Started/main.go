package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func main() {
	//rdb初始化
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "845924",
		DB:       0,
	})
	//创建上下文
	ctx := context.Background()
	err := rdb.Set(ctx, "goredistest", "goredistestvalue", time.Second*10).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "goredistest").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

	// 如果希望用原生命令 可以使用Do方法
	res, err := rdb.Do(ctx, "get", "goredistest").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
