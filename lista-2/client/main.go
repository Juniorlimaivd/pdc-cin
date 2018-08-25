package main

import (
	"bufio"
	"fmt"
)

type transferData struct {
	PayerID string
	PayeeID string
	Amount  float32
}

type accOperation struct {
	AccID  string
	Amount float32
}

func transferCommand(rw *bufio.ReadWriter, reader *bufio.Reader) error {
	testData := transferData{
		PayerID: "AC1",
		PayeeID: "AC2",
		Amount:  200}

	sendString(rw, "TRANSFER")
	sendEncondedData(rw, testData)

	return nil
}

func getBalanceCommand(rw *bufio.ReadWriter, reader *bufio.Reader) error {
	sendString(rw, "BALANCE")

	fmt.Print("Account ID: ")
	id, _ := readString(reader)

	sendString(rw, id)

	balance, _ := recvData(rw)

	fmt.Println("Balance: " + balance)

	return nil
}

func withdrawCommand(rw *bufio.ReadWriter, reader *bufio.Reader) error {
	sendString(rw, "WITHDRAW")

	fmt.Print("Account ID: ")
	id, _ := readString(reader)

	fmt.Print("Amount: ")
	amount, _ := readFloat32(reader)

	testData := accOperation{
		AccID:  id,
		Amount: amount}

	sendEncondedData(rw, testData)

	return nil
}

func depositCommand(rw *bufio.ReadWriter, reader *bufio.Reader) error {
	sendString(rw, "DEPOSIT")
	fmt.Print("Account ID: ")
	id, _ := readString(reader)

	fmt.Print("Amount: ")
	amount, _ := readFloat32(reader)

	testData := accOperation{
		AccID:  id,
		Amount: amount}

	sendEncondedData(rw, testData)

	return nil
}

func main() {

	client := NewClient()
	client.AddCommandFunc(CommandInfo{
		shortName:   "T",
		longName:    "Transfer Command",
		description: "Do somenthing"}, transferCommand)

	client.AddCommandFunc(CommandInfo{
		shortName:   "B",
		longName:    "Get Balance Command",
		description: "Do somenthing"}, getBalanceCommand)

	client.AddCommandFunc(CommandInfo{
		shortName:   "W",
		longName:    "Withdraw Command",
		description: "Do somenthing"}, withdrawCommand)

	client.AddCommandFunc(CommandInfo{
		shortName:   "D",
		longName:    "Deposit Command",
		description: "Do something"}, depositCommand)

	err := client.Start("127.0.0.1:8081")

	if err != nil {
		fmt.Println("Couldn't connect to server")
		return
	}
}
