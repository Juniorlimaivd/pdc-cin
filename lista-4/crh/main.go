package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/tealeg/xlsx"
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

func unPacketToString(data []byte) string {
	var result string
	var buf bytes.Buffer
	buf.Write(data)
	decoder := gob.NewDecoder(&buf)

	decoder.Decode(&result)

	return result
}

func main() {
	times := 10
	filename := "crh_middleware.xlsx"

	currentFile := xlsx.NewFile()
	sheet, _ := currentFile.AddSheet("Sheet1")

	crh := newClientRequestHandler("localhost", 12345, "middleware")

	crh.connect()

	accInfo := AccountInformation{ID: "1234"}
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
