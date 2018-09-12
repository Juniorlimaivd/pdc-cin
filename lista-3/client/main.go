package main

import (
	"fmt"
	"log"
	"net/rpc"
)

const serverAddress = "localhost"

func main() {
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := 10.10
	var reply string
	err = client.Call("Account.Deposit", args, &reply)
	if err != nil {
		log.Fatal("account error:", err)
	}
	fmt.Printf("Reply: %s\n", reply)

}
