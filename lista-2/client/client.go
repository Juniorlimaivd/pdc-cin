package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/pkg/errors"
)

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
	acc.handleCommands(conn)
	return nil
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
