package main

import (
	"context"
	"log"
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

	//3. 声明队列
	q, err := ch.QueueDeclare(
		"hello", //队列名
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Cannt declare a queue: %v", err)
	}

	//4. 准备要发布的消息
	body := "Hello RabbitMQ"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	//5. 发布消息到队列
	err = ch.PublishWithContext(
		ctx,
		"",
		q.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		log.Fatalf("Cannt publish a message: %v", err)
	}
	log.Printf(" [x] Sent %s", body)
}
