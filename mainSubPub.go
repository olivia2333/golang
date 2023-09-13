package main

import (
	"imooc/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("newProduct")
	for i := 0; i < 100; i++ {
		rabbitmq.PublishPub("No." + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}
