package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

type handleFunc func(*bufio.ReadWriter)

type Endpoint struct {
	listener net.Listener
	handler  map[string]handleFunc
	m        sync.RWMutex
}

func newEndpoint() *Endpoint {
	return &Endpoint{
		handler: map[string]handleFunc{},
	}
}

func (e *Endpoint) addHandleFunc(name string, f handleFunc) {
	e.m.Lock()
	e.handler[name] = f
	e.m.Unlock()
}

func (e *Endpoint) listen(port string) error {
	var err error
	e.listener, err = net.Listen("tcp", port)
	if err != nil {
		return errors.Wrapf(err, "Unable to listen on port %s\n", port)
	}
	log.Println("Listen on", e.listener.Addr().String())
	for {
		conn, err := e.listener.Accept()
		log.Println("Accept a connection request from", conn.RemoteAddr())
		if err != nil {
			log.Println("Failed accepting a connection request:", err)
			continue
		}
		go e.handleMessages(conn)
	}
}

func (e *Endpoint) handleMessages(conn net.Conn) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	defer conn.Close()

	for {
		log.Println("Handle incoming commands...")
		cmd, err := e.parseCommand(rw)

		switch {
		case err == io.EOF:
			log.Println(conn.RemoteAddr(), "close this connection.\n   ---")
			return
		case err != nil:
			log.Println("\nError reading command. Got: '"+cmd+"'\n", err)
			return
		case cmd == "":
			log.Println("Ignoring empty command")
			continue
		}

		log.Print("Receive command '" + cmd + "'")

		handleCommand := e.getCommandHandler(cmd)
		if handleCommand != nil {
			handleCommand(rw)
		}
	}
}

func (e *Endpoint) parseCommand(rw *bufio.ReadWriter) (string, error) {
	cmd, err := rw.ReadString('\n')

	if err != nil {
		return "", err
	}

	cmd = strings.Trim(cmd, "\n ")
	return cmd, nil
}

func (e *Endpoint) getCommandHandler(cmd string) handleFunc {
	e.m.RLock()
	handleCommand, ok := e.handler[cmd]
	e.m.RUnlock()
	if !ok {
		log.Println("Command '" + cmd + "' is not registered.")
		return nil
	}
	return handleCommand
}
