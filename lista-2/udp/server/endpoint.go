package main

import (
	"bytes"
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

type handleFunc func(*Endpoint, *RequestOperationData, net.Addr)

// Endpoint handles all incoming data fro socket and redirect to proper handler
type Endpoint struct {
	listener net.PacketConn
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
	e.listener, err = net.ListenPacket("udp", port)
	if err != nil {
		return errors.Wrapf(err, "Unable to listen on port %s\n", port)
	}

	log.Println("Listen on", e.listener.LocalAddr().String())

	buffer := make([]byte, 1024)

	for {
		n, addr, err := e.listener.ReadFrom(buffer)

		if err != nil {
			log.Println("Failed receiving message:", err)
		}

		go e.handleMessage(buffer, n, addr)
	}
}

func (e *Endpoint) handleMessage(buffer []byte, n int, addr net.Addr) {

	request, err := e.parseCommand(buffer, n)

	switch {
	case err == io.EOF:
		log.Println("close this connection.\n   ---")
		return
	case err != nil:
		log.Println("\nError reading command. Got: \n", err)
		return
	case request.OperationType == "":
		log.Println("Ignoring empty command, closing connection...")
		return
	}

	log.Print("Receive command '" + request.OperationType + "'")

	handleCommand := e.getCommandHandler(request)
	if handleCommand != nil {
		handleCommand(e, &request, addr)
	}

}

func (e *Endpoint) parseCommand(buffer []byte, n int) (RequestOperationData, error) {
	var request RequestOperationData

	reader := bytes.NewReader(buffer)
	dec := gob.NewDecoder(reader)
	dec.Decode(&request)

	return request, nil
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

func (e *Endpoint) sendResultDescription(result string, addr net.Addr) {
	pkt := OperationResult{ResultDescription: result}
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)

	encoder.Encode(pkt)

	e.listener.WriteTo(buffer.Bytes(), addr)

}
