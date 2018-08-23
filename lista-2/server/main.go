package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
)

var commChannel = make(chan TransferData)

var accounts = map[string]*Account{
	"AC1": &Account{balance: 1000},
	"AC2": &Account{balance: 2000}}

func handleTransfer(rw *bufio.ReadWriter) {
	log.Println("Handling Transfer")
	var transferData TransferData
	dec := gob.NewDecoder(rw)
	err := dec.Decode(&transferData)
	if err != nil {
		log.Println("Error decoding TransferData:", err)
		return
	}
	commChannel <- transferData

}

func handleGetBalance(rw *bufio.ReadWriter) {
	id, err := rw.ReadString('\n')
	if err != nil {
		log.Println("Error reading ID from buffer:", err)
		return
	}
	fmt.Println("Account id:", id)
}

func handleWithdraw(rw *bufio.ReadWriter) {

}

func handleDeposit(rw *bufio.ReadWriter) {

}

func transferWorker() {
	log.Println("Transfer Worker Started")
	for {
		transferData := <-commChannel
		log.Printf("Transfering $%.2f from %s to %s...\n", transferData.Amount, transferData.PayerID, transferData.PayeeID)

		payer := accounts[transferData.PayerID]
		payee := accounts[transferData.PayeeID]
		payer.withdraw(transferData.Amount)
		payee.deposit(transferData.Amount)
		log.Println("- Successful Transaction -")
		log.Printf("After transaction:")
		log.Printf("%s: $%.2f\n", transferData.PayerID, payer.balance)
		log.Printf("%s: $%.2f\n", transferData.PayeeID, payee.balance)

	}
}

func main() {
	go transferWorker()

	endpoint := NewEndpoint()
	endpoint.AddHandleFunc("TRANSFER", handleTransfer)
	endpoint.AddHandleFunc("BALANCE", handleGetBalance)
	endpoint.AddHandleFunc("WITHDRAW", handleWithdraw)
	endpoint.AddHandleFunc("DEPOSIT", handleDeposit)
	endpoint.Listen(":8081")

	fmt.Scanln()
}
