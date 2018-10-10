package server

import "log"

// AccountsManager handles all accounts
type AccountsManager struct {
	Accs map[string]*Account
}

// Account handles information about one account
type Account struct {
	Balance float64
}

func (acc *Account) deposit(amount float64) error {
	acc.Balance += amount
	return nil
}

func (acc *Account) withdraw(amount float64) error {
	acc.Balance -= amount
	return nil
}

// GetBalance returns the balance of the requested account
func (accMngr *AccountsManager) GetBalance(accID string) float64 {
	log.Printf("Getting balance of %s account", accID)
	acc := accMngr.Accs[accID]

	return acc.Balance
}

// Deposit puts the money in the account
func (accMngr *AccountsManager) Deposit(accID string, amount float64) string {
	log.Printf("Depositing $ %f into %s account", amount, accID)
	acc := accMngr.Accs[accID]

	acc.deposit(amount)

	return "amount successfully deposited"
}

// Withdraw draws the money from the account
func (accMngr *AccountsManager) Withdraw(accID string, amount float64) string {
	log.Printf("Withdrawing $ %f from %s account", amount, accID)
	acc := accMngr.Accs[accID]

	acc.withdraw(amount)

	return "amount successfully withdrawn"
}

// Transfer withdraw money from the first account and puts into the other account
func (accMngr *AccountsManager) Transfer(payerID string, payeeID string, amount float64) string {
	log.Printf("Transfering $%.2f from %s to %s...\n", amount, payerID, payeeID)

	payer := accMngr.Accs[payerID]
	payee := accMngr.Accs[payeeID]

	payer.withdraw(amount)
	payee.deposit(amount)

	return "successful transfer"
}
