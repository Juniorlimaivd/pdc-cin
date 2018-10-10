package server

import (
	"strconv"

	"../global"
)

// RunTest ...
func RunTest() {
	accs := make(map[string]*Account)
	initialBalance := 1000.0
	for i := 0; i < global.AccsNumber; i++ {
		accID := global.AccInitial + strconv.Itoa(i+1)
		accs[accID] = &Account{Balance: initialBalance}
	}
	accManager := AccountsManager{Accs: accs}
	invoker := NewInvoker(&accManager)
	invoker.Invoke()

}
