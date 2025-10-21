package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func main() {
	// SET依旧跟别处一样 天然去重
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "845924",
		DB:       0,
	})
	ctx := context.Background()

	// SADD
	rdb.SAdd(ctx, "setkey", 100)
	rdb.SAdd(ctx, "setkey", 100)
	rdb.SAdd(ctx, "setkey", 201)
	rdb.SAdd(ctx, "setkey", 21)

	//SCARD
	size, _ := rdb.SCard(ctx, "setkey").Result()
	fmt.Println(size)

	// SIsmember
	// 判断是否在列表里面

	// SRem
	// 删除方法
	// rdb.SRem(ctx, "setkey", 100)

	// Spop
	// 随机返回集合中的元素 并且删除返回的元素
	val1, _ := rdb.SPop(ctx, "setkey").Result()
	val2, _ := rdb.SPopN(ctx, "setkey", 2).Result()
	fmt.Println(val1)
	fmt.Println(val2)
}
