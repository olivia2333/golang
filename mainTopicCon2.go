package main

import "imooc/RabbitMQ"

func main() {
	imoocOne := RabbitMQ.NewRabbitMQTopic("exImoocTopic", "imooc.*.two")
	imoocOne.ReceiveTopic()
}
