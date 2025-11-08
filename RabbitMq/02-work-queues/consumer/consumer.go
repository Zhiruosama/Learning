package main

import (
	"bytes"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	//1. 连接到RabbitMQ
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

	//3. 声明队列 队列名要和生产者保持一致
	q, err := ch.QueueDeclare("task_queue", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Cannt declare a queue: %v", err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Cannt register a consumer: %v", err)
	}

	log.Printf(" [*] 等待任务中。按 CTRL+C 退出")

	for d := range msgs {
		log.Printf(" [x] 收到 %s", d.Body)

		dotCount := bytes.Count(d.Body, []byte("."))
		t := time.Duration(dotCount)
		time.Sleep(t * time.Second)

		log.Printf("Task complete")
	}
}
