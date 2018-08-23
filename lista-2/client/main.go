package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
)

type transferData struct {
	PayerID string
	PayeeID string
	Amount  float32
}

func transferCommand(rw *bufio.ReadWriter, reader *bufio.Reader) error {
	testData := transferData{
		PayerID: "AC1",
		PayeeID: "AC2",
		Amount:  200}

	enc := gob.NewEncoder(rw)

	rw.WriteString("TRANSFER\n")
	err := rw.Flush()
	if err != nil {
		fmt.Println("Flush failed")
		return err
	}

	err = enc.Encode(testData)
	if err != nil {
		fmt.Println("Error encoding", err)
		return err
	}
	err = rw.Flush()
	if err != nil {
		fmt.Println("Flush failed")
		return err
	}

	return nil
}

func main() {

	client := NewClient()
	client.AddCommandFunc(CommandInfo{
		shortName:   "T",
		longName:    "Transfer Command",
		description: "Do somenthing"}, transferCommand)

	err := client.Start("127.0.0.1:8081")

	if err != nil {
		fmt.Println("Couldn't connect to server")
		return
	}
}
