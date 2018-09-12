package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/rpc"
	"strconv"
	"time"

	"github.com/tealeg/xlsx"
)

type OpFunc func()

var commands = map[string]OpFunc{
	"B": getBalance,
	"W": withdraw,
	"D": deposit,
	"T": transfer,
}

var accsNumber = 4

var currentFile *xlsx.File

var sheet *xlsx.Sheet

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
	//log.Printf("Getting balance from %s account", accID)

	var reply float32
	start := time.Now()
	err = client.Call("AccountsManager.GetBalance", accID, &reply)
	end := time.Now()

	row := sheet.AddRow()
	cell := row.AddCell()
	cell.SetFloat(float64(end.Sub(start).Nanoseconds()) / 1000000.) // in miliseconds

	if err != nil {
		log.Fatal("account error:", err)
	}
<<<<<<< HEAD
	log.Printf("Server reply: %.2f", reply)
=======
	//log.Printf("Server reply: %f.2", reply)
>>>>>>> fdcb1d16b855e0392ac7284f59b47fab1beb0efd
}

func deposit() {
	accOpArgs := randomAccOpArgs()
<<<<<<< HEAD
	log.Printf("Depositing $%.2f into %s account", accOpArgs.Amount, accOpArgs.AccID)
=======
	//log.Printf("Depositing $ %f into %s account", accOpArgs.Amount, accOpArgs.AccID)
>>>>>>> fdcb1d16b855e0392ac7284f59b47fab1beb0efd

	var depositReply string
	start := time.Now()
	err = client.Call("AccountsManager.Deposit", accOpArgs, &depositReply)
	end := time.Now()

	row := sheet.AddRow()
	cell := row.AddCell()
	cell.SetFloat(float64(end.Sub(start).Nanoseconds()) / 1000000.) // in miliseconds

	if err != nil {
		log.Fatal("account error:", err)
	}
	//log.Printf("Server reply: %s", depositReply)

}

func withdraw() {
	accOpArgs := randomAccOpArgs()
<<<<<<< HEAD
	log.Printf("Withdrawing $%.2f from %s account", accOpArgs.Amount, accOpArgs.AccID)
=======
	//log.Printf("Withdrawing $ %f from %s account", accOpArgs.Amount, accOpArgs.AccID)
>>>>>>> fdcb1d16b855e0392ac7284f59b47fab1beb0efd

	var withdrawReply string
	start := time.Now()
	err = client.Call("AccountsManager.Withdraw", accOpArgs, &withdrawReply)
	end := time.Now()
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.SetFloat(float64(end.Sub(start).Nanoseconds()) / 1000000.) // in miliseconds

	if err != nil {
		log.Fatal("account error:", err)
	}
	//log.Printf("Server reply: %s", withdrawReply)

}

func transfer() {
	transferArgs := randomTransferArgs()
	//log.Printf("Transfering $%.2f from %s to %s...\n", transferArgs.Amount, transferArgs.PayerID, transferArgs.PayeeID)

	var transferReply string

	start := time.Now()
	err = client.Call("AccountsManager.Transfer", transferArgs, &transferReply)
	end := time.Now()
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.SetFloat(float64(end.Sub(start).Nanoseconds()) / 1000000.) // in miliseconds

	if err != nil {
		log.Fatal("account error:", err)
	}
	//log.Printf("Server reply: %s", transferReply)
}

func main() {

	filenames := [4]string{"tcp_balance_10000.xlsx",
		"tcp_withdraw_10000.xlsx",
		"tcp_deposit_10000.xlsx",
		"tcp_transfer_10000.xlsx"}
	i := 0
	if err != nil {
		log.Fatal("dialing:", err)
	}
<<<<<<< HEAD
	for index := 0; index < 1000000; index++ {
		for _, command := range commands {
=======

	for _, command := range commands {
		currentFile = xlsx.NewFile()
		sheet, _ = currentFile.AddSheet("Sheet1")
		for index := 0; index < 10000; index++ {
			//log.Printf("Call %s", key)
>>>>>>> fdcb1d16b855e0392ac7284f59b47fab1beb0efd
			command()
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		}

		currentFile.Save(filenames[i])
		fmt.Println("Finished: ", filenames[i])
		i++

	}

}
