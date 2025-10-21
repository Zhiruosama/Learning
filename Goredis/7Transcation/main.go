package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func main() {
	// redis的事务有些特殊
	// 不具备原子性 事务中某一个命令出错
	// 其他命令也会继续执行
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "845924",
		DB:       0,
	})
	ctx := context.Background()

	// 开启一个TxPipeline事务
	pipe := rdb.TxPipeline()
	// 执行事务操作 可以通过Pipe读写redis
	incr := pipe.Incr(ctx, "tx_pipeline_coutner")
	pipe.Expire(ctx, "tx_pipeline_coutner", time.Hour)
	// 通过exec函数提交redis事务
	pipe.Exec(ctx)
	// 提交后 查询事务操作的记过
	fmt.Println(incr.Val())
}
