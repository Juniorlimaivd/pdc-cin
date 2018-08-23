package main

type TransferData struct {
	PayerID string
	PayeeID string
	Amount  float32
}

type Account struct {
	balance float32
}

func (acc *Account) deposit(amount float32) {
	acc.balance += amount
}

func (acc *Account) withdraw(amount float32) {
	acc.balance -= amount
}
