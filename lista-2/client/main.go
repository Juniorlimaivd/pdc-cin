package main

import (
	"bufio"
	"fmt"
)

func transferCommand(rw *bufio.ReadWriter, reader *bufio.Reader) error {
	fmt.Print(" * payer ID: ")
	payerID, _ := readString(reader)
	fmt.Print(" * payee ID: ")
	payeeID, _ := readString(reader)
	fmt.Print(" * amount: ")
	amount, _ := readFloat32(reader)

	transferData := TransferData{
		PayerID: payerID,
		PayeeID: payeeID,
		Amount:  amount}

	sendString(rw, "TRANSFER")
	sendEncondedData(rw, transferData)
	check, _ := recvString(rw)

	if check == "OK" {
		fmt.Print("\n % Sucessful operation %\n\n")
	}
	return nil
}

func getBalanceCommand(rw *bufio.ReadWriter, reader *bufio.Reader) error {
	fmt.Print(" * account ID: ")
	id, _ := readString(reader)

	sendString(rw, "BALANCE")
	sendString(rw, id)

	balance, _ := recvString(rw)
	fmt.Printf(" ---------------------\n + balance: %s\n\n", balance)

	return nil
}

func withdrawCommand(rw *bufio.ReadWriter, reader *bufio.Reader) error {
	fmt.Print(" * account ID: ")
	id, _ := readString(reader)
	fmt.Print(" * amount: ")
	amount, _ := readFloat32(reader)

	accOperation := AccOperation{
		AccID:  id,
		Amount: amount}

	sendString(rw, "WITHDRAW")
	sendEncondedData(rw, accOperation)

	return nil
}

func depositCommand(rw *bufio.ReadWriter, reader *bufio.Reader) error {
	fmt.Print("Account ID: ")
	id, _ := readString(reader)
	fmt.Print("Amount: ")
	amount, _ := readFloat32(reader)

	accOperation := AccOperation{
		AccID:  id,
		Amount: amount}

	sendString(rw, "DEPOSIT")
	sendEncondedData(rw, accOperation)

	return nil
}

func main() {
	client := NewClient()
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
