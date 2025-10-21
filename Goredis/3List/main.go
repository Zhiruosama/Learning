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

	// LPUSH
	// 从左侧添加数据
	rdb.LPush(ctx, "key", "data1")
	rdb.LPush(ctx, "key", 1, 2, 3, 4, 5, 6)

	// RPOP
	// 从右侧弹出数据
	rdb.RPop(ctx, "key")

	// RPUSH
	// LPOP
	// 同上操作

	// LRange 取出数据
	vals, _ := rdb.LRange(ctx, "key", 0, -1).Result()
	fmt.Println(vals)

	// LRem 从左往右删找到的第一个目标值
	rdb.LRem(ctx, "key", 1, 5)
	// count输入多少就根据顺序删除多少个目标值
	// 如果打算从右往左删除的话就将count设置为负数
	// 0的话就是全部删除
}
