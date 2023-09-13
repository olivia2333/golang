package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// amqp://username:password@rabbit_address:port/virtualhost
const MQURL = "amqp://imoocuser:imoocuser@127.0.0.1:5672/imooc"

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//queue
	QueueName string
	//exchange
	Exchange string
	//key
	Key string
	//connection
	Mqurl string
}

// create rabbitmq instance
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{QueueName: queueName,
		Exchange: exchange, Key: key, Mqurl: MQURL}
	var err error

	// create rabbitmq connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "connection creation fail")

	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "fail to get channel")
	return rabbitmq
}

// close channel and connection
func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

// error handling
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// simple mode 1: create simple rabbitmq instance
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	// use default exchange
	rabbitmq := NewRabbitMQ(queueName, "", "")
	return rabbitmq
}

// simple mode 2: create
func (r *RabbitMQ) PublishSimple(message string) {
	// 1. acquire channel, if channel not exist will create automatically, else will skip
	//     guarantee existence of channel
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		// if messages durable
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	// 2. send message to channel
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		// if cannot find exchange and mandatory == true, messages will be sent back to sender
		false,
		// if immediate == true, and find channel doesn't have receiver, will send back messages to seder
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
}

func (r *RabbitMQ) ConsumeSimple() {
	// 1. acquire channel, if channel not exist will create automatically, else will skip
	//     guarantee existence of channel
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		// if messages durable
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	// 2. receive msgs
	msgs, err := r.channel.Consume(
		r.QueueName,
		// differentiate multiple consumers
		"",
		true,
		false,
		// cannot use messages in the same connection to other consumers
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	// couroutines on messages
	go func() {
		for d := range msgs {
			// deal messages logic
			log.Printf("Receive a message: %s", d.Body)
		}
	}()
	log.Printf("[*] Waiting for messages, To exit press CTRL+C")
	<-forever
}

// publish/subscribe
func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchangeName, "")
	var err error

	// create rabbitmq connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "connection creation fail")

	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "fail to get channel")
	return rabbitmq
}

// create
func (r *RabbitMQ) PublishPub(message string) {
	// 1. tries to create exchange
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		// cannot be used for client to send message, only used to bind exchanges
		false,
		false,
		nil,
	)

	r.failOnErr(err, "Failed to declare an exchange")

	// 2. send msgs
	err = r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// publish/subscribe consumer
func (r *RabbitMQ) ReceiveSub() {
	// 1. tries to create exchange
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		// cannot be used for client to send message, only used to bind exchanges
		false,
		false,
		nil,
	)

	r.failOnErr(err, "Failed to declare an exchange")

	// 2. tries to create channel
	q, err := r.channel.QueueDeclare(
		"", // random create channel
		false,
		false,
		true, // true
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declared a queue")

	// bind channel to exchange
	err = r.channel.QueueBind(
		q.Name,
		"",
		r.Exchange,
		false,
		nil,
	)

	// consume messages
	msgs, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	// couroutines on messages
	go func() {
		for d := range msgs {
			// deal messages logic
			log.Printf("Receive a message: %s", d.Body)
		}
	}()
	log.Printf("[*] Waiting for messages, To exit press CTRL+C")
	<-forever
}

// routing
// publish/subscribe
func NewRabbitMQRouting(exchangeName string, routingKey string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchangeName, routingKey)
	var err error

	// create rabbitmq connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "connection creation fail")

	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "fail to get channel")
	return rabbitmq
}

// create
func (r *RabbitMQ) PublishRouting(message string) {
	// 1. tries to create exchange
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		// cannot be used for client to send message, only used to bind exchanges
		false,
		false,
		nil,
	)

	r.failOnErr(err, "Failed to declare an exchange")

	// 2. send msgs
	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// routing consumer
func (r *RabbitMQ) ReceiveRouting() {
	// 1. tries to create exchange
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		// cannot be used for client to send message, only used to bind exchanges
		false,
		false,
		nil,
	)

	r.failOnErr(err, "Failed to declare an exchange")

	// 2. tries to create channel
	q, err := r.channel.QueueDeclare(
		"", // random create channel
		false,
		false,
		true, // true
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declared a queue")

	// bind channel to exchange
	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil,
	)

	// consume messages
	msgs, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	// couroutines on messages
	go func() {
		for d := range msgs {
			// deal messages logic
			log.Printf("Receive a message: %s", d.Body)
		}
	}()
	log.Printf("[*] Waiting for messages, To exit press CTRL+C")
	<-forever
}

// topic
func NewRabbitMQTopic(exchangeName string, routingKey string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchangeName, routingKey)
	var err error

	// create rabbitmq connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "connection creation fail")

	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "fail to get channel")
	return rabbitmq
}

// create
func (r *RabbitMQ) PublishTopic(message string) {
	// 1. tries to create exchange
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic",
		true,
		false,
		// cannot be used for client to send message, only used to bind exchanges
		false,
		false,
		nil,
	)

	r.failOnErr(err, "Failed to declare an exchange")

	// 2. send msgs
	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// topic consumer
// e.g. imooc.* -> imooc.hello, but not imooc.hello.one, need imooc.#
func (r *RabbitMQ) ReceiveTopic() {
	// 1. tries to create exchange
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic",
		true,
		false,
		// cannot be used for client to send message, only used to bind exchanges
		false,
		false,
		nil,
	)

	r.failOnErr(err, "Failed to declare an exchange")

	// 2. tries to create channel
	q, err := r.channel.QueueDeclare(
		"", // random create channel
		false,
		false,
		true, // true
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declared a queue")

	// bind channel to exchange
	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil,
	)

	// consume messages
	msgs, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	// couroutines on messages
	go func() {
		for d := range msgs {
			// deal messages logic
			log.Printf("Receive a message: %s", d.Body)
		}
	}()
	log.Printf("[*] Waiting for messages, To exit press CTRL+C")
	<-forever
}
