package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"log"
	"net"
	"os"
	"strconv"
	"time"

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

	accInfo := AccountInformation{ID: "1234"}

	// connect to this socket
	conn, _ := net.Dial("udp", "127.0.0.1:8081")
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)

	for i := 0; i < times; i++ {
		data := packetData(accInfo)

		start := time.Now()
		enc := gob.NewEncoder(writer)
		enc.Encode(data)
		writer.Flush()

		// listen for reply
		var data2 []byte
		decoder := gob.NewDecoder(reader)

		decoder.Decode(&data2)

		end := time.Now()
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.SetFloat(float64(end.Sub(start).Nanoseconds()) / 1000000.) // in miliseconds

		dec := unPacketToString(data2)
		if dec != "OK" {
			log.Fatal("Some error has occurred")
		}
	}

	currentFile.Save(filename)
}
