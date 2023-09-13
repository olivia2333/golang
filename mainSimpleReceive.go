package main

import "imooc/RabbitMQ"

func main() {
	// note: to run work mode, just launch two receivers
	rabbitmq := RabbitMQ.NewRabbitMQSimple("imoocSimple")
	rabbitmq.ConsumeSimple()
}
