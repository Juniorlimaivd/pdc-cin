package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/tealeg/xlsx"
)

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

func main() {
	if len(os.Args) != 4 {
		log.Fatal("Invalid number of arguments")
	}

	handlerType := os.Args[1]
	times, err := strconv.Atoi(os.Args[2])
	failOnError(err, "Failed to #times of execution")

	filename := os.Args[3]
	if exists(filename) {
		log.Fatal("File \"" + filename + "\" already exists")
	}

	currentFile := xlsx.NewFile()
	sheet, _ := currentFile.AddSheet("Sheet1")

	crh := newClientRequestHandler("localhost", 12345, handlerType)

	err = crh.connect()
	failOnError(err, "Failed to connect to Server")

	accInfo := AccountInformation{ID: "1234"}
	log.Println("Sending " + strconv.Itoa(times) + " requests to " + handlerType + " request handler")
	log.Println("Logging time spent into " + filename + " file")

	for i := 0; i < times; i++ {

		data := packetData(accInfo)

		start := time.Now()

		crh.send(data)
		pkt := crh.receive()

		end := time.Now()

		row := sheet.AddRow()
		cell := row.AddCell()
		cell.SetFloat(float64(end.Sub(start).Nanoseconds()) / 1000000.) // in miliseconds

		//fmt.Println("Operation took ", end.Sub(start))

		result := unPacketToString(pkt)

		if result == "OK" {
			fmt.Println("Successful operation")
		}
	}

	currentFile.Save(filename)

}
