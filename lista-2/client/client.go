package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/tealeg/xlsx"
)

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

type CommandFunc func(*bufio.ReadWriter, *bufio.Reader) error

type CommandInfo struct {
	shortName   string
	longName    string
	description string
}

type Client struct {
	command     map[string]CommandFunc
	commandInfo map[string]CommandInfo
}

func NewClient() *Client {
	return &Client{
		command:     map[string]CommandFunc{},
		commandInfo: map[string]CommandInfo{},
	}
}

func (acc *Client) AddCommandFunc(commandInfo CommandInfo, f CommandFunc) {
	key := commandInfo.shortName
	acc.command[key] = f
	acc.commandInfo[key] = commandInfo
}

func (acc *Client) Start(addr string) error {
	log.Println("Dial " + addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return errors.Wrap(err, "Dialing "+addr+" failed")
	}
	acc.showDescriptions()
	//acc.handleCommands(conn)
	acc.performOperationTest(conn, "T", 1000)
	return nil
}

func (acc *Client) performOperationTest(conn net.Conn, operation string, times int) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	reader := bufio.NewReader(os.Stdin)

	defer conn.Close()

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	cmd := operation

	for i := 0; i < times; i++ {
		fmt.Println("Performing tests...")

		handleCommand := acc.getCommandHandler(cmd)
		if handleCommand != nil {
			log.Println("Performing - ", acc.commandInfo[cmd].longName, " -")

			start := time.Now()
			err := handleCommand(rw, reader)
			end := time.Now()

			fmt.Println("Operation ", i, " took ", end.Sub(start).Seconds())

			row = sheet.AddRow()
			cell = row.AddCell()
			cell.SetFloat(end.Sub(start).Seconds() * 1000)

			time.Sleep(100 * time.Millisecond)
			if err != nil {
				log.Print(cmd, "Failed")
			}
		}
	}

	err = file.Save("tcp_transfer.xlsx")

	if err != nil {
		fmt.Printf(err.Error())
	}
}

func (acc *Client) handleCommands(conn net.Conn) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	reader := bufio.NewReader(os.Stdin)

	defer conn.Close()

	for {
		fmt.Print("Type a command: ")

		cmd, _ := acc.parseCommand(reader)

		handleCommand := acc.getCommandHandler(cmd)
		if handleCommand != nil {
			log.Print("Performing - ", acc.commandInfo[cmd].longName, " -")
			err := handleCommand(rw, reader)
			if err != nil {
				log.Print(cmd, "Failed")
			}
		}
	}

}

func (acc *Client) parseCommand(reader *bufio.Reader) (string, error) {
	cmd, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	cmd = strings.Trim(cmd, "\n ")
	return cmd, nil
}

func (acc *Client) getCommandHandler(cmd string) CommandFunc {
	handleCommand, ok := acc.command[cmd]

	if !ok {
		log.Println("Command '" + cmd + "' is not registered.")
		return nil
	}
	return handleCommand
}

func (acc *Client) showDescriptions() {
	fmt.Println("\nCOMMAND LIST:")
	for _, v := range acc.commandInfo {
		fmt.Printf(" ** %s (%s) - %s\n", v.shortName, v.longName, v.description)
	}
	fmt.Println()
}
