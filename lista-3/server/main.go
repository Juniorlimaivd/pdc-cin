package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

var accounts = map[string]*Account{
	"AC1": &Account{Balance: 1000},
	"AC2": &Account{Balance: 2000},
	"AC3": &Account{Balance: 3000},
	"AC4": &Account{Balance: 4000}}

type AccountsManager struct {
	Accs map[string]*Account
}

type Account struct {
	Balance float32
}

type AccOpArgs struct {
	AccID  string
	Amount float32
}

type TransferArgs struct {
	PayerID string
	PayeeID string
	Amount  float32
}

func (acc *Account) deposit(amount float32) error {
	acc.Balance += amount
	return nil
}

func (acc *Account) withdraw(amount float32) error {
	acc.Balance -= amount
	return nil
}

func (accMngr *AccountsManager) GetBalance(accID string, reply *float32) error {
	log.Printf("Getting balance from of %s account", accID)
	acc := accMngr.Accs[accID]

	*reply = acc.Balance
	return nil
}

func (accMngr *AccountsManager) Deposit(args AccOpArgs, reply *string) error {
	log.Printf("Depositing $ %f into %s account", args.Amount, args.AccID)
	acc := accMngr.Accs[args.AccID]

	acc.deposit(args.Amount)

	*reply = "amount successfully deposited"
	return nil
}

func (accMngr *AccountsManager) Withdraw(args AccOpArgs, reply *string) error {
	log.Printf("Withdrawing $ %f on %s account", args.Amount, args.AccID)
	acc := accMngr.Accs[args.AccID]

	acc.withdraw(args.Amount)

	*reply = "amount successfully withdrawn"
	return nil
}

func (accMngr *AccountsManager) Transfer(args *TransferArgs, reply *string) error {
	log.Printf("Transfering $%.2f from %s to %s...\n", args.Amount, args.PayerID, args.PayeeID)

	payer := accMngr.Accs[args.PayerID]
	payee := accMngr.Accs[args.PayeeID]

	payer.withdraw(args.Amount)
	payee.deposit(args.Amount)

	*reply = "successful transfer"
	return nil
}

func main() {
	accManager := new(AccountsManager)
	accManager.Accs = accounts
	rpc.Register(accManager)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
	fmt.Scanln()

}
