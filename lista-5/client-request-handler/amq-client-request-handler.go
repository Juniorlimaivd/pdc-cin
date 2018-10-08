package main

import (
	"github.com/streadway/amqp"
)

// AMQClientRequestHandler handles tcp connections
type AMQClientRequestHandler struct {
	host         string
	port         int
	id           string
	ch           *amqp.Channel
	inFromServer amqp.Queue
	outToServer  amqp.Queue
	msgs         <-chan amqp.Delivery
}

func newAMQClientRequestHandler(host string, port int) *AMQClientRequestHandler {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	//failOnError(err, "Failed to connect to RabbitMQ")

	ch, _ := conn.Channel()
	//failOnError(err, "Failed to open a channel")

	inFromServer, _ := ch.QueueDeclare(
		"inFromServer", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	msgs, _ := ch.Consume(
		inFromServer.Name, // queue
		"",                // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	//failOnError(err, "Failed to register a consumer")

	outToServer, _ := ch.QueueDeclare(
		"outToServer", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	//failOnError(err, "Failed to declare a queue")

	return &AMQClientRequestHandler{
		host:         host,
		port:         port,
		ch:           ch,
		inFromServer: inFromServer,
		outToServer:  outToServer,
		msgs:         msgs,
	}
}

func (c *AMQClientRequestHandler) connect() error {

	return nil
}

func (c *AMQClientRequestHandler) send(data []byte) error {
	c.ch.Publish(
		"",                 // exchange
		c.outToServer.Name, // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})
	//failOnError(err, "Failed to publish a message")

	return nil
}

func (c *AMQClientRequestHandler) receive() []byte {
	data := <-c.msgs
	return data.Body
}
