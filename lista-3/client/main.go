package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type AccOpArgs struct {
	AccID  string
	Amount float32
}

type TransferArgs struct {
	PayerID string
	PayeeID string
	Amount  float32
}

const serverAddress = "localhost"

func main() {
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call

	// Get Balance
	accID := "AC1"
	var reply float32
	err = client.Call("AccountsManager.GetBalance", accID, &reply)
	if err != nil {
		log.Fatal("account error:", err)
	}
	fmt.Printf("Balance: %f\n", reply)

	// Deposit
	accOpArgs := AccOpArgs{"AC1", 32}
	var depositReply string
	err = client.Call("AccountsManager.Deposit", accOpArgs, &depositReply)
	if err != nil {
		log.Fatal("account error:", err)
	}
	fmt.Printf("Deposit reply: %s\n", depositReply)

	// Withdraw
	var withdrawReply string
	err = client.Call("AccountsManager.Withdraw", accOpArgs, &withdrawReply)
	if err != nil {
		log.Fatal("account error:", err)
	}
	fmt.Printf("Deposit reply: %s\n", withdrawReply)

	// Transfer
	transferArgs := TransferArgs{"AC1", "AC2", 32}
	var transferReply string
	err = client.Call("AccountsManager.Transfer", transferArgs, &transferReply)
	if err != nil {
		log.Fatal("account error:", err)
	}
	fmt.Printf("Deposit reply: %s\n", transferReply)

}
