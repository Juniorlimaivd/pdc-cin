package main

import (
	"flag"
	"log"

	"./client"
	"./server"
)

func createServer() {
	log.Println("handling create server")
	server.RunTest()
}

func main() {
	mwType := flag.String(
		"type",
		"",
		"Describes the middleware type to be initialized\n* Available options\n- client\n- server")

	filename := flag.String("output", "", "Name of the output file")
	reqName := flag.String("request", "", "Request type\n* Available options\n- getBalance\n- deposit\n- withdraw\n- transfer")

	times := flag.Int("times", 0, "")
	flag.Parse()

	switch *mwType {
	case "server":
		createServer()
	case "client":
		client.RunTest(*filename, *times, *reqName)
	default:
		log.Fatalf("Unknown middleware type: unable to find \"%s\"", *mwType)
	}
}
