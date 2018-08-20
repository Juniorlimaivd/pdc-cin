package main

import (
	"fmt"
)

// Account is cool
type Account struct {
	id      int64
	balance float64
}

// TransferData is cool
type TransferData struct {
	payer  *Account
	payee  *Account
	amount float64
}

func (acc *Account) deposit(amount float64) {
	acc.balance += amount
}

func (acc *Account) withdraw(amount float64) {
	acc.balance -= amount
}

func transfer(payer *Account, payee *Account, amount float64, comm chan TransferData) {
	comm <- TransferData{payer: payer, payee: payee, amount: amount}
}

func transferWorker(comm chan TransferData) {
	for {
		transferData := <-comm
		fmt.Println("Transfering from", transferData.payer.id, "to", transferData.payee.id, "...")
		transferData.payer.withdraw(transferData.amount)
		transferData.payee.deposit(transferData.amount)
	}
}

func main() {
	fmt.Println("olar")
	commChannel := make(chan TransferData)

	go transferWorker(commChannel)

	ac1 := Account{id: 01,
		balance: 109020}

	ac2 := Account{id: 02,
		balance: 993280}

	go func() {
		for {
			transfer(&ac1, &ac2, 10, commChannel)
		}
	}()

	go func() {
		for {
			transfer(&ac2, &ac1, 10, commChannel)
		}
	}()

	fmt.Scanln()
}
