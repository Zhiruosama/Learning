package main

import (
	"context"
	"fmt"

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

	// ZADD添加数据方法
	// 这里特殊的地方就在每个key会有一个score
	// 并且会被自动排序
	rdb.ZAdd(ctx, "zkey", redis.Z{Score: 2.5, Member: "zhangsan"})
	rdb.ZAdd(ctx, "zkey", redis.Z{Score: 5, Member: "wangwu"})
	rdb.ZAdd(ctx, "zkey", redis.Z{Score: 4, Member: "lisi"})

	// ZCount
	// 统计某个分数范围内的元素个数
	// 分数左闭右闭
	size, _ := rdb.ZCount(ctx, "zkey", "1", "3").Result()
	fmt.Println(size)

	// ZincrBy
	// 自增
	rdb.ZIncrBy(ctx, "zkey", 1, "wangwu")

	// ZRangeBy
	// 根据范围进行查找
	op := redis.ZRangeBy{
		Min:    "1",
		Max:    "10",
		Offset: 0,
		Count:  3,
	}
	vals, _ := rdb.ZRangeByScoreWithScores(ctx, "zkey", &op).Result()
	for _, val := range vals {
		fmt.Println(val.Member)
		fmt.Println(val.Score)
	}
}
