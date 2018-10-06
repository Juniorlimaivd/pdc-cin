package serverrequesthandler

import (
	"github.com/streadway/amqp"
)

type AMQServerRequestHandler struct {
	port        string
	msgs        <-chan amqp.Delivery
	ch          *amqp.Channel
	outToClient *amqp.Queue
}

func NewAMQServerRequestHandler(port string) *AMQServerRequestHandler {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	inFromClient, err := ch.QueueDeclare(
		"outToServer", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	msgs, err := ch.Consume(
		inFromClient.Name, // queue
		"",                // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	failOnError(err, "Failed to register a consumer")

	outToClient, err := ch.QueueDeclare(
		"inFromServer", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return &AMQServerRequestHandler{
		msgs:        msgs,
		port:        port,
		ch:          ch,
		outToClient: &outToClient,
	}
}

func (c *AMQServerRequestHandler) send(data []byte) {
	err := c.ch.Publish(
		"",                 // exchange
		c.outToClient.Name, // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})
	failOnError(err, "Failed to publish a message")
}

func (c *AMQServerRequestHandler) receive() []byte {
	data := <-c.msgs
	return data.Body
}
