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

func transfer(ac1 *Account, ac2 *Account, amount float64) {
	fmt.Println("Transfering from", ac1.id, "to ", ac2.id)

	fmt.Println("Locking ", ac1.id)
	ac1.acessControl.Lock()
	fmt.Println("Locking ", ac2.id)
	ac2.acessControl.Lock()

	fmt.Println("Transfering...")
	ac1.withdraw(amount)
	ac2.deposit(amount)

	fmt.Println("Unlocking ", ac1.id)
	ac1.acessControl.Unlock()
	fmt.Println("Unlocking ", ac2.id)
	ac2.acessControl.Unlock()
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
