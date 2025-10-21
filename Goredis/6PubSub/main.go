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

	// 订阅
	// channel_*这种通配符的方法也可以
	sub := rdb.Subscribe(ctx, "channel1")
	// for ch := range sub.Channel() {
	// 	fmt.Println(ch.Channel)
	// 	fmt.Println(ch.Payload)
	// }
	for {
		msg, _ := sub.ReceiveMessage(ctx)
		fmt.Println(msg.Channel)
		fmt.Println(msg.Payload)
	}
}
