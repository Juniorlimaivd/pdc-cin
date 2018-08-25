package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"strings"
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
	log.Println(" -- Balance --")
	id, err := rw.ReadString('\n')
	id = strings.Trim(id, "\n ")
	if err != nil {
		log.Println("Error reading ID from buffer:", err)
		return
	}
	log.Printf("%s: $%.2f\n", id, accounts[id].balance)
	rw.WriteString(fmt.Sprintf("%.2f\n", accounts[id].balance))
	if err != nil {
		log.Println("Cannot write to connection.\n", err)
	}
	rw.Flush()
}

func handleWithdraw(rw *bufio.ReadWriter) {
	log.Println(" -- Withdraw --")

	var accOperation AccOperation
	dec := gob.NewDecoder(rw)
	err := dec.Decode(&accOperation)
	if err != nil {
		log.Println("Error decoding accOperation:", err)
		return
	}
	log.Println(accOperation)
}

func handleDeposit(rw *bufio.ReadWriter) {
	log.Println(" -- Deposit --")

	var accOperation AccOperation
	dec := gob.NewDecoder(rw)
	err := dec.Decode(&accOperation)
	if err != nil {
		log.Println("Error decoding accOperation:", err)
		return
	}
	log.Println(accOperation)
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

	endpoint := newEndpoint()
	endpoint.addHandleFunc("TRANSFER", handleTransfer)
	endpoint.addHandleFunc("BALANCE", handleGetBalance)
	endpoint.addHandleFunc("WITHDRAW", handleWithdraw)
	endpoint.addHandleFunc("DEPOSIT", handleDeposit)
	endpoint.listen(":8081")

	fmt.Scanln()
}
