package main

import (
	"bufio"
	"log"
	"strings"
)

func sendString(rw *bufio.ReadWriter, data string) error {
	rw.WriteString(data + "\n")
	err := rw.Flush()
	if err != nil {
		log.Println("Flush failed")
		return err
	}
	return nil
}

func recvString(rw *bufio.ReadWriter) (string, error) {
	data, err := rw.ReadString('\n')
	return strings.Trim(data, "\n "), err
}
