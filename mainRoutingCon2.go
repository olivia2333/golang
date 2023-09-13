package main

import "imooc/RabbitMQ"

func main() {
	imoocTwo := RabbitMQ.NewRabbitMQRouting("exImooc", "imooc_two")
	imoocTwo.ReceiveRouting()
}
