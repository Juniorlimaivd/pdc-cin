package main

// TransferData is a packet for transfer request
type TransferData struct {
	PayerID string
	PayeeID string
	Amount  float32
}

// OperationResult encapsulates a packet for server response
type OperationResult struct {
	ResultDescription string
}

// AccOperation is a packet for withdraw or deposit request
type AccOperation struct {
	AccID  string
	Amount float32
}

// AccountInformation is a packet for balance request
type AccountInformation struct {
	Id string
}

// Account is a struct for handle balance
type Account struct {
	balance float32
}

func (acc *Account) deposit(amount float32) {
	acc.balance += amount
}

func (acc *Account) withdraw(amount float32) {
	acc.balance -= amount
}
