package main

import (
	"bytes"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("RabbitMQ创建失败: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("管道创建失败: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("queue_ttl_demo", false, true, false, false, nil)
	if err != nil {
		log.Fatalf("队列创建失败: %v", err)
	}

	// 此时必须为手动 ack
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("无法注册一个消费者: %v", err)
	}

	log.Printf(" [*] 等待任务(Qos=1, 公平分发)。按 CTRL+C 退出")

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)

		dotCount := bytes.Count(d.Body, []byte("."))
		t := time.Duration(dotCount)
		time.Sleep(t * time.Second)

		log.Printf("任务完成")

		d.Ack(false)
	}
}
