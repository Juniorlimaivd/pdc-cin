package main

import (
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

	listener, err := net.ListenPacket("udp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err.Error())
		// return errors.Wrapf(err, "Unable to listen on port %s\n", port)
	}
	fmt.Println("Listen on", listener.LocalAddr().String())

	buffer := make([]byte, 1024)

	for {
		_, clientAddr, err := listener.ReadFrom(buffer)

		if err != nil {
			fmt.Println(err.Error())
		}

		var data []byte

		reader := bytes.NewReader(buffer)
		dec := gob.NewDecoder(reader)
		err = dec.Decode(&data)

		unPacketToAccInfo(data)
		// fmt.Println("Message Received:" + msg_rcv.ID)

		msg := "OK"
		msg_enc := packetData(msg)

		var buffer bytes.Buffer

		encoder := gob.NewEncoder(&buffer)

		encoder.Encode(msg_enc)

		listener.WriteTo(buffer.Bytes(), clientAddr)

	}
}
