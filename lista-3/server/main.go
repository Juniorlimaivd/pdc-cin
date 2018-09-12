package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Account struct {
	Id      string
	Balance float32
}

func (acc *Account) Deposit(amount float32, reply *string) error {
	acc.Balance += amount
	*reply = "amount successfully deposited"
	return nil
}

func (acc *Account) Withdraw(amount float32, reply *string) error {
	acc.Balance -= amount
	*reply = "amount successfully withdrawn"
	return nil
}

func main() {
	account := new(Account)
	rpc.Register(account)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
	fmt.Scanln()

}
