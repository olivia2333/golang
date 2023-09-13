package main

import (
	"fmt"
	"imooc/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("imoocSimple")
	rabbitmq.PublishSimple("Hello World")
	fmt.Println("sent success")
}
