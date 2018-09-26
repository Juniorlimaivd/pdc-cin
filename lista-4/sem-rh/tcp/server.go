package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
)

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

func unPacketToAccInfo(data []byte) AccountInformation {
	var result AccountInformation
	var buf bytes.Buffer
	buf.Write(data)
	decoder := gob.NewDecoder(&buf)

	decoder.Decode(&result)

	return result
}

func main() {
	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)

	for {
		// will listen for message to process ending in newline (\n)
		decoder := gob.NewDecoder(reader)
		var data2 []byte
		decoder.Decode(&data2)
		if data2 == nil {
			break
		}
		dec := unPacketToAccInfo(data2)
		// output message received
		fmt.Println("Message Received:" + dec.ID)
		// sample process for string received
		newmessage := "OK"
		// send new string back to client
		data := packetData(newmessage)
		enc := gob.NewEncoder(writer)
		enc.Encode(data)
		writer.Flush()
	}
}
