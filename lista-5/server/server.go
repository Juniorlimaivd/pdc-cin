package main

import (
	"log"

	"./middleware"
)

// AccountsManager handles all accounts
type AccountsManager struct {
	Accs map[string]*Account
}

// Account handles information about one account
type Account struct {
	Balance float32
}

func (acc *Account) deposit(amount float32) error {
	acc.Balance += amount
	return nil
}

func (acc *Account) withdraw(amount float32) error {
	acc.Balance -= amount
	return nil
}

// GetBalance returns the balance of the requested account
func (accMngr *AccountsManager) GetBalance(accID string) float32 {
	log.Printf("Getting balance of %s account", accID)
	acc := accMngr.Accs[accID]

	return acc.Balance
}

// Deposit puts the money in the account
func (accMngr *AccountsManager) Deposit(accID string, amount float32) string {
	//log.Printf("Depositing $ %f into %s account", args.Amount, args.AccID)
	acc := accMngr.Accs[accID]

	acc.deposit(amount)

	return "amount successfully deposited"
}

// Withdraw draws the money from the account
func (accMngr *AccountsManager) Withdraw(accID string, amount float32) string {
	//log.Printf("Withdrawing $ %f from %s account", args.Amount, args.AccID)
	acc := accMngr.Accs[accID]

	acc.withdraw(amount)

	return "amount successfully withdrawn"
}

// Transfer withdraw money from the first account and puts into the other account
func (accMngr *AccountsManager) Transfer(payerID string, payeeID string, amount float32) string {
	//log.Printf("Transfering $%.2f from %s to %s...\n", args.Amount, args.PayerID, args.PayeeID)

	payer := accMngr.Accs[payerID]
	payee := accMngr.Accs[payeeID]

	payer.withdraw(amount)
	payee.deposit(amount)

	return "successful transfer"
}

func main() {
	accManager := AccountsManager{Accs: map[string]*Account{
		"AC1": &Account{Balance: 1000},
		"AC2": &Account{Balance: 2000},
		"AC3": &Account{Balance: 3000},
		"AC4": &Account{Balance: 4000}}}

	invoker := middleware.NewInvoker(&accManager)

	invoker.Invoke()

}
