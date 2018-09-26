package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// AccountInformation holds info about request account
type AccountInformation struct {
	ID string
}

func packetData(data interface{}) []byte {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)

	encoder.Encode(data)
	return buf.Bytes()
}

func unPacketToAccInfo(data []byte) AccountInformation {
	var result AccountInformation
	var buf bytes.Buffer
	buf.Write(data)
	decoder := gob.NewDecoder(&buf)

	decoder.Decode(&result)

	return result
}

func main() {

	if len(os.Args) != 2 {
		log.Fatal("Invalid number of arguments")
	}
	handlerType := os.Args[1]
	port := 12345

	srh, err := newServerRequestHandler(handlerType, port)
	failOnError(err, "Failed to create a new server request handler")

	for {
		data := srh.receive()
		if data == nil {
			break
		}
		accInfo := unPacketToAccInfo(data)
		log.Println(accInfo.ID)
		sentString := "OK"
		pkt := packetData(sentString)
		srh.send(pkt)
	}
}
