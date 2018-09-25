package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// TransferInfo is cool
type TransferInfo struct {
	Transfer         TransferData
	TransferEndpoint *Endpoint
}

var commChannel = make(chan TransferInfo)

var accounts = map[string]*Account{
	"AC1": &Account{balance: 1000},
	"AC2": &Account{balance: 2000},
	"AC3": &Account{balance: 3000},
	"AC4": &Account{balance: 4000}}

func handleTransfer(e *Endpoint, r *RequestOperationData) {
	//log.Println("Handling Transfer")

	var transferData TransferData

	reader := bytes.NewReader(r.Data)
	dec := gob.NewDecoder(reader)
	err := dec.Decode(&transferData)

	if err != nil {
		log.Println("Error decoding TransferData:", err)
		return
	}

	transferInfo := TransferInfo{Transfer: transferData, TransferEndpoint: e}
	commChannel <- transferInfo

}

func handleGetBalance(e *Endpoint, r *RequestOperationData) {
	//log.Println(" -- Balance --")

	var accountInformation AccountInformation

	reader := bytes.NewReader(r.Data)
	dec := gob.NewDecoder(reader)
	err := dec.Decode(&accountInformation)

	if err != nil {
		log.Println("Error decoding Balance:", err)
		return
	}

	account, ok := accounts[accountInformation.Id]

	var result float32

	if !ok {
		result = 0
	} else {
		result = account.balance
	}

	//log.Printf("%s: $%.2f\n", accountInformation.Id, result)

	e.sendResultDescription(fmt.Sprintf("%.2f", result))

}

func handleWithdraw(e *Endpoint, r *RequestOperationData) {
	//log.Println(" -- Withdraw --")

	var accOperation AccOperation
	reader := bytes.NewReader(r.Data)

	dec := gob.NewDecoder(reader)
	err := dec.Decode(&accOperation)

	if err != nil {
		log.Println("Error decoding accOperation:", err)
		return

	}

	payee, ok := accounts[accOperation.AccID]

	if !ok {
		accounts[accOperation.AccID] = &Account{balance: 0}
		payee = accounts[accOperation.AccID]
	}

	payee.withdraw(accOperation.Amount)

	//log.Println(accOperation)

	e.sendResultDescription("OK")
}

func handleDeposit(e *Endpoint, r *RequestOperationData) {
	//log.Println(" -- Deposit --")

	var accOperation AccOperation

	reader := bytes.NewReader(r.Data)

	dec := gob.NewDecoder(reader)
	err := dec.Decode(&accOperation)

	if err != nil {
		log.Println("Error decoding accOperation:", err)
		return
	}

	payer, ok := accounts[accOperation.AccID]

	if !ok {
		accounts[accOperation.AccID] = &Account{balance: 0}
		payer = accounts[accOperation.AccID]
	}

	payer.deposit(accOperation.Amount)

	//log.Println(accOperation)

	e.sendResultDescription("OK")
}

func transferWorker() {
	log.Println("Transfer Worker Started")
	for {
		transferInfo := <-commChannel
		transferData := transferInfo.Transfer
		//log.Printf("Transfering $%.2f from %s to %s...\n", transferData.Amount, transferData.PayerID, transferData.PayeeID)

		payer, ok := accounts[transferData.PayerID]

		if !ok {
			accounts[transferData.PayerID] = &Account{balance: 0}
			payer = accounts[transferData.PayerID]
		}

		payee, ok := accounts[transferData.PayeeID]

		if !ok {
			accounts[transferData.PayeeID] = &Account{balance: 0}
			payee = accounts[transferData.PayeeID]
		}

		payer.withdraw(transferData.Amount)
		payee.deposit(transferData.Amount)
		//log.Println("- Successful Transaction -")
		//log.Printf("After transaction:")
		//log.Printf("%s: $%.2f\n", transferData.PayerID, payer.balance)
		//log.Printf("%s: $%.2f\n", transferData.PayeeID, payee.balance)
		transferInfo.TransferEndpoint.sendResultDescription("OK")
	}
}

func main() {
	go transferWorker()

	endpoint := newEndpoint()
	endpoint.addHandleFunc("TRANSFER", handleTransfer)
	endpoint.addHandleFunc("BALANCE", handleGetBalance)
	endpoint.addHandleFunc("WITHDRAW", handleWithdraw)
	endpoint.addHandleFunc("DEPOSIT", handleDeposit)
	endpoint.listen(":12345")

	fmt.Scanln()
}
