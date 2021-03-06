package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"strconv"
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

func sendEncondedData(rw *bufio.ReadWriter, data interface{}) error {
	enc := gob.NewEncoder(rw)
	err := enc.Encode(data)
	if err != nil {
		fmt.Println("Error encoding", err)
		return err
	}
	err = rw.Flush()
	if err != nil {
		fmt.Println("Flush failed")
		return err
	}
	return nil
}

func recvOperationResult(rw *bufio.ReadWriter) (OperationResult, error) {
	var result OperationResult

	decoder := gob.NewDecoder(rw)

	decoder.Decode(&result)

	return result, nil
}

func recvString(rw *bufio.ReadWriter) (string, error) {
	data, err := rw.ReadString('\n')
	return strings.Trim(data, "\n "), err
}

func readString(reader *bufio.Reader) (string, error) {
	data, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	data = strings.Trim(data, "\n")
	return data, nil
}

func readFloat32(reader *bufio.Reader) (float32, error) {
	data, _ := readString(reader)
	value, err := strconv.ParseFloat(data, 32)
	if err != nil {
		fmt.Println("Error parsing float: ", err)
	}
	return float32(value), nil
}

func packtRequestData(opType string, data interface{}) RequestOperationData {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	encoder.Encode(data)
	newData := buf.Bytes()

	return RequestOperationData{OperationType: opType, Data: newData}
}
