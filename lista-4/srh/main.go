package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

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

	port := 12345

	srh := newServerRequestHandler("middleware", port)

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
		log.Printf("[x] Sent %s", sentString)

	}
}
