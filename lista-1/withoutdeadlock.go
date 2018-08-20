package main

import (
	"fmt"
	"sync"
)

// Account is cool
type Account struct {
	id           int64
	balance      float64
	acessControl sync.Mutex
}

func (acc *Account) deposit(amount float64) {
	acc.balance += amount
}

func (acc *Account) withdraw(amount float64) {
	acc.balance -= amount
}

func transfer(payer *Account, payee *Account, amount float64) {
	fmt.Println("Transfering from", payer.id, "to ", payee.id)

	if payer.id < payee.id {
		fmt.Println("Locking ", payer.id)
		payer.acessControl.Lock()
		fmt.Println("Locking ", payee.id)
		payee.acessControl.Lock()
	} else {
		fmt.Println("Locking ", payee.id)
		payee.acessControl.Lock()
		fmt.Println("Locking ", payer.id)
		payer.acessControl.Lock()
	}

	fmt.Println("Transfering...")
	payer.withdraw(amount)
	payee.deposit(amount)

	if payer.id < payee.id {
		fmt.Println("Unlocking ", payer.id)
		payer.acessControl.Unlock()
		fmt.Println("Unlocking ", payee.id)
		payee.acessControl.Unlock()
	} else {
		fmt.Println("Unlocking ", payee.id)
		payee.acessControl.Unlock()
		fmt.Println("Unlocking ", payer.id)
		payer.acessControl.Unlock()
	}

}

func main() {
	fmt.Println("olar")

	ac1 := Account{id: 01,
		balance: 109020}

	ac2 := Account{id: 02,
		balance: 993280}

	go func() {
		for {
			transfer(&ac1, &ac2, 10)
		}
	}()

	go func() {
		for {
			transfer(&ac2, &ac1, 10)
		}
	}()

	fmt.Scanln()
}
