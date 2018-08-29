package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"flag"
	"fmt"
	"math/rand"
	"strconv"
)

// RequestOperationData is cool
type RequestOperationData struct {
	OperationType string
	Data          []byte
}

func transferCommand(rw *bufio.ReadWriter, reader *bufio.Reader, testMode bool) error {

	var payerID, payeeID string
	var amount float32

	if !testMode {
		fmt.Print(" * payer ID: ")
		payerID, _ = readString(reader)
		fmt.Print(" * payee ID: ")
		payeeID, _ = readString(reader)
		fmt.Print(" * amount: ")
		amount, _ = readFloat32(reader)
	} else {
		fmt.Print(" * payer ID: ")
		payerID = strconv.Itoa(rand.Intn(100))
		fmt.Print(payerID)
		fmt.Print("\n * payee ID: ")
		payeeID = strconv.Itoa(rand.Intn(100))
		fmt.Print(payeeID)
		fmt.Print("\n * amount: ")
		amount = rand.Float32() * 100.0
		fmt.Print(amount)
		fmt.Print("\n")
	}

	transferData := TransferData{
		PayerID: payerID,
		PayeeID: payeeID,
		Amount:  amount}

	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)

	encoder.Encode(transferData)
	data := buf.Bytes()

	request := RequestOperationData{OperationType: "TRANSFER", Data: data}

	sendEncondedData(rw, request)

	check, _ := recvOperationResult(rw)

	if check.ResultDescription == "OK" {
		fmt.Print("\n % Sucessful operation %\n\n")
	}
	return nil
}

func getBalanceCommand(rw *bufio.ReadWriter, reader *bufio.Reader, testMode bool) error {
	var id string

	if !testMode {
		fmt.Print(" * account ID: ")
		id, _ = readString(reader)
	} else {
		fmt.Print(" * account ID: ")
		id = strconv.Itoa(rand.Intn(100))
	}

	accData := AccountInformation{Id: id}

	requestPkt := packtRequestData("BALANCE", accData)

	sendEncondedData(rw, requestPkt)

	result, _ := recvOperationResult(rw)
	fmt.Printf(" ---------------------\n + balance: %s\n\n", result.ResultDescription)

	return nil
}

func withdrawCommand(rw *bufio.ReadWriter, reader *bufio.Reader, testMode bool) error {

	var id string
	var amount float32

	if !testMode {
		fmt.Print(" * account ID: ")
		id, _ = readString(reader)
		fmt.Print(" * amount: ")
		amount, _ = readFloat32(reader)
	} else {
		fmt.Print(" * account ID: ")
		id = strconv.Itoa(rand.Intn(100))
		fmt.Print(id + "\n")
		fmt.Print(" * amount: ")
		amount = rand.Float32() * 100.0
		fmt.Print(amount, "\n")
	}

	accOperation := AccOperation{
		AccID:  id,
		Amount: amount}

	requestPkt := packtRequestData("WITHDRAW", accOperation)
	sendEncondedData(rw, requestPkt)

	check, _ := recvOperationResult(rw)

	if check.ResultDescription == "OK" {
		fmt.Print("\n % Sucessful operation %\n\n")
	}

	return nil
}

func depositCommand(rw *bufio.ReadWriter, reader *bufio.Reader, testMode bool) error {
	var id string
	var amount float32

	if !testMode {
		fmt.Print(" * account ID: ")
		id, _ = readString(reader)
		fmt.Print(" * amount: ")
		amount, _ = readFloat32(reader)
	} else {
		fmt.Print(" * account ID: ")
		id = strconv.Itoa(rand.Intn(100))
		fmt.Print(id + "\n")
		fmt.Print(" * amount: ")
		amount = rand.Float32() * 100.0
		fmt.Print(amount, "\n")
	}

	accOperation := AccOperation{
		AccID:  id,
		Amount: amount}

	requestPkt := packtRequestData("DEPOSIT", accOperation)
	sendEncondedData(rw, requestPkt)

	check, _ := recvOperationResult(rw)

	if check.ResultDescription == "OK" {
		fmt.Print("\n % Sucessful operation %\n\n")
	}

	return nil
}

func main() {
	testBool := flag.Bool("test", false, "true for use testmode")

	flag.Parse()

	client := NewClient(*testBool)
	client.AddCommandFunc(CommandInfo{
		shortName:   "T",
		longName:    "Transfer Command",
		description: "Transfer <amount> from <payer ID> to <payee ID>"}, transferCommand)

	client.AddCommandFunc(CommandInfo{
		shortName:   "B",
		longName:    "Get Balance Command",
		description: "Get balance of <account ID>"}, getBalanceCommand)

	client.AddCommandFunc(CommandInfo{
		shortName:   "W",
		longName:    "Withdraw Command",
		description: "Withdraw from the <account ID>"}, withdrawCommand)

	client.AddCommandFunc(CommandInfo{
		shortName:   "D",
		longName:    "Deposit Command",
		description: "Deposit into the <account ID>"}, depositCommand)

	err := client.Start("127.0.0.1:8081")

	if err != nil {
		fmt.Println("Couldn't connect to server")
		return
	}
}
