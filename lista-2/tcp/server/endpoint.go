package main

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
	"net"
	"sync"

	"github.com/pkg/errors"
)

// RequestOperationData holds information about sended data
type RequestOperationData struct {
	OperationType string
	Data          []byte
}

type handleFunc func(*Endpoint, *RequestOperationData)

// Endpoint handles all incoming data fro socket and redirect to proper handler
type Endpoint struct {
	listener net.Listener
	rw       *bufio.ReadWriter
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
	e.rw = bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	defer conn.Close()

	for {
		//log.Println("Handling incoming commands...")
		request, err := e.parseCommand(e.rw)

		switch {
		case err == io.EOF:
			log.Println(conn.RemoteAddr(), "close this connection.\n   ---")
			return
		case err != nil:
			log.Println("\nError reading command. Got: \n", err)
			return
		case request.OperationType == "":
			log.Println("Ignoring empty command, closing connection...")
			return
		}

		//log.Print("Receive command '" + request.OperationType + "'")

		handleCommand := e.getCommandHandler(request)
		if handleCommand != nil {
			handleCommand(e, &request)
		}
	}
}

func (e *Endpoint) parseCommand(rw *bufio.ReadWriter) (RequestOperationData, error) {
	var request RequestOperationData

	decoder := gob.NewDecoder(rw)

	err := decoder.Decode(&request)

	return request, err
}

func (e *Endpoint) getCommandHandler(request RequestOperationData) handleFunc {
	e.m.RLock()
	handleCommand, ok := e.handler[request.OperationType]
	e.m.RUnlock()
	if !ok {
		log.Println("Command '" + request.OperationType + "' is not registered.")
		return nil
	}
	return handleCommand
}

func (e *Endpoint) sendResultDescription(result string) {
	pkt := OperationResult{ResultDescription: result}

	encoder := gob.NewEncoder(e.rw)

	encoder.Encode(pkt)

	e.rw.Flush()

}
