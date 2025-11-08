package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	//1. 连接rabbitmq
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Cannt connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	//2. 创建一个通道
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Cannt open a channel: %v", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare("logs", "fanout", false, true, false, false, nil)
	if err != nil {
		log.Fatalf("无法声明 exchange: %v", err)
	}

	//4. 准备要发布的消息
	body := bodyFrom(os.Args)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	//5. 发布消息到队列
	err = ch.PublishWithContext(ctx, "logs", "", false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	if err != nil {
		log.Fatalf("Cannt publish a message: %v", err)
	}
	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
