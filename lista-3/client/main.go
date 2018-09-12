package main

import (
	"log"
	"math/rand"
	"net/rpc"
	"strconv"
)

type OpFunc func()

var commands = map[string]OpFunc{
	"B": getBalance,
	"W": withdraw,
	"D": deposit,
	"T": transfer,
}

var accsNumber = 4

var client, err = rpc.DialHTTP("tcp", serverAddress+":1234")

const serverAddress = "localhost"

func randomAccID() string {
	return "AC" + strconv.Itoa(rand.Intn(accsNumber)+1)
}

func randomAmount() float32 {
	return rand.Float32() * 100.0
}

func randomAccOpArgs() AccOpArgs {
	accID := randomAccID()
	amount := randomAmount()
	return AccOpArgs{accID, amount}
}

func randomTransferArgs() TransferArgs {
	payerID := randomAccID()
	payeeID := randomAccID()
	amount := randomAmount()
	return TransferArgs{payerID, payeeID, amount}
}

func getBalance() {
	accID := randomAccID()
	log.Printf("Getting balance from %s account", accID)

	var reply float32
	err = client.Call("AccountsManager.GetBalance", accID, &reply)
	if err != nil {
		log.Fatal("account error:", err)
	}
	log.Printf("Server reply: %f.2", reply)
}

func deposit() {
	accOpArgs := randomAccOpArgs()
	log.Printf("Depositing $ %f into %s account", accOpArgs.Amount, accOpArgs.AccID)

	var depositReply string
	err = client.Call("AccountsManager.Deposit", accOpArgs, &depositReply)
	if err != nil {
		log.Fatal("account error:", err)
	}
	log.Printf("Server reply: %s", depositReply)

}

func withdraw() {
	accOpArgs := randomAccOpArgs()
	log.Printf("Withdrawing $ %f from %s account", accOpArgs.Amount, accOpArgs.AccID)

	var withdrawReply string
	err = client.Call("AccountsManager.Withdraw", accOpArgs, &withdrawReply)
	if err != nil {
		log.Fatal("account error:", err)
	}
	log.Printf("Server reply: %s", withdrawReply)

}

func transfer() {
	transferArgs := randomTransferArgs()
	log.Printf("Transfering $%.2f from %s to %s...\n", transferArgs.Amount, transferArgs.PayerID, transferArgs.PayeeID)

	var transferReply string
	err = client.Call("AccountsManager.Transfer", transferArgs, &transferReply)
	if err != nil {
		log.Fatal("account error:", err)
	}
	log.Printf("Server reply: %s", transferReply)
}

func main() {

	if err != nil {
		log.Fatal("dialing:", err)
	}
	for index := 0; index < 10; index++ {
		for key, command := range commands {
			log.Printf("Call %s", key)
			command()
		}
	}

}
