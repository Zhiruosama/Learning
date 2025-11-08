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

	//3. 声明direct类型的exchange
	err = ch.ExchangeDeclare("logs_topic", "topic", false, true, false, false, nil)
	if err != nil {
		log.Fatalf("无法声明exchange: %v", err)
	}

	//4. 从命令行获取日志级别和消息
	// 用法: go run main.go error "数据库连接失败"
	severity := severityFrom(os.Args)
	body := bodyFrom(os.Args)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	//5. 发布消息
	err = ch.PublishWithContext(ctx, "logs_topic", severity, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	if err != nil {
		log.Fatalf("发布消息失败: %v", err)
	}

	log.Printf(" [x] 发送 [%s] 日志: %s", severity, body)
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

func severityFrom(args []string) string {
	// os.Args[0] 是程序名本身
	// 我们检查是否有第二个参数 (os.Args[1])，即用户输入的第一条实际参数
	if len(args) < 2 {
		// 如果没有提供参数，默认返回 "info"
		return "info"
	}
	// 如果提供了参数，返回第一个参数作为 severity (路由键)
	return args[1]
}
