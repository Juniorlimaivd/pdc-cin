package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/streadway/amqp"
	"github.com/tealeg/xlsx"
)

// exists returns whether the given file or directory exists or not
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// AccountInformation is cool
type AccountInformation struct {
	ID string
}

func packetData(data interface{}) []byte {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)

	encoder.Encode(data)
	return buf.Bytes()
}

func unPacketToString(data []byte) string {
	var result string
	var buf bytes.Buffer
	buf.Write(data)
	decoder := gob.NewDecoder(&buf)

	decoder.Decode(&result)

	return result
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Invalid number of arguments")
	}

	times, err := strconv.Atoi(os.Args[1])
	failOnError(err, "Failed to #times of execution")

	filename := os.Args[2]
	if exists(filename) {
		log.Fatal("File \"" + filename + "\" already exists")
	}

	currentFile := xlsx.NewFile()
	sheet, _ := currentFile.AddSheet("Sheet1")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"sendToServer", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare a queue")

	rq, errq := ch.QueueDeclare(
		"receiveFromServer", // name
		false,               // durable
		false,               // delete when usused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	failOnError(errq, "Failed to declare a queue")

	msgs, err := ch.Consume(
		rq.Name, // queue
		"",      // consumer
		true,    // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
	failOnError(err, "Failed to register a consumer")

	for i := 0; i < times; i++ {
		accInfo := AccountInformation{ID: "1234"}
		data := packetData(accInfo)
		start := time.Now()
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        data,
			})
		//failOnError(err, "Failed to publish a message")

		result := <-msgs
		end := time.Now()

		row := sheet.AddRow()
		cell := row.AddCell()
		cell.SetFloat(float64(end.Sub(start).Nanoseconds()) / 1000000.) // in miliseconds

		response := unPacketToString(result.Body)
		if response != "OK" {
			log.Fatal("Some error has occurred")
		}
		//fmt.Println(response)
	}

	currentFile.Save(filename)

}
