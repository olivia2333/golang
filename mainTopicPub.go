package main

import (
	"fmt"
	"imooc/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	imoocTopicOne := RabbitMQ.NewRabbitMQTopic("exImoocTopic", "imooc.topic.one")
	imoocTopicTwo := RabbitMQ.NewRabbitMQTopic("exImoocTopic", "imooc.topic.two")

	for i := 0; i <= 10; i++ {
		imoocTopicOne.PublishTopic("Hello imooc topic one!" + strconv.Itoa(i))
		imoocTopicTwo.PublishTopic("Hello imooc topic two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
