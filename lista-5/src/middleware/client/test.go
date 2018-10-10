package client

import (
	"log"
	"os"
	"time"

	"../global"
	"github.com/tealeg/xlsx"
)

type requestMethod func(proxy *Proxy)

var requestsMethod = map[string]requestMethod{
	"getBalance": randomGetBalance,
	"deposit":    randomDeposit,
	"withdraw":   randomWithdraw,
	"transfer":   randomTransfer,
}

// exists returns whether the given file or directory exists or not
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func randomGetBalance(proxy *Proxy) {
	proxy.getBalance(global.RandomAccID())
}

func randomDeposit(proxy *Proxy) {
	proxy.deposit(global.RandomAccID(), global.RandomAmount())
}

func randomWithdraw(proxy *Proxy) {
	proxy.withdraw(global.RandomAccID(), global.RandomAmount())
}

func randomTransfer(proxy *Proxy) {
	proxy.transfer(global.RandomAccID(), global.RandomAccID(), global.RandomAmount())
}

// RunTest ...
func RunTest(filename string, times int, reqName string) {
	log.Println("Running test")
	if exists(filename) {
		log.Fatalf("File \"%s\" already exists\n", filename)
	}
	reqMethod := requestsMethod[reqName]
	if reqMethod == nil {
		log.Fatalf("Unknown request method: \"%s\" was not found\n", reqName)
	}
	if times <= 0 {
		log.Fatalln("Invalid number of executions")
	}
	proxy := newProxy("localhost", 1234)
	currentFile := xlsx.NewFile()
	sheet, _ := currentFile.AddSheet("Sheet1")
	log.Printf("Running %d times \"%s\" request\n", times, reqName)
	log.Printf("Logging time spent into %s file\n", filename)
	for i := 0; i < times; i++ {
		start := time.Now()
		reqMethod(proxy)
		end := time.Now()
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.SetFloat(float64(end.Sub(start).Nanoseconds()) / 1000000.) // in miliseconds
	}
	currentFile.Save(filename)
}
