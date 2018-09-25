package main

import (
	"bufio"
	"encoding/gob"
	"log"
	"net"
)

type TCPServerRequestHandler struct {
	port        int
	listener    net.Listener
	rw          *bufio.ReadWriter
	outToClient *bufio.Reader
	inToClient  *bufio.Writer
}

func NewTCPServerRequestHandler(port string) *TCPServerRequestHandler {
	tcpSRH := new(TCPServerRequestHandler)
	var err error
	tcpSRH.listener, err = net.Listen("tcp", port)
	if err != nil {
		// return errors.Wrapf(err, "Unable to listen on port %s\n", port)
	}
	log.Println("Listen on", tcpSRH.listener.Addr().String())
	conn, err := tcpSRH.listener.Accept()
	log.Println("Accept a connection request from", conn.RemoteAddr())

	tcpSRH.inToClient = bufio.NewWriter(conn)
	tcpSRH.outToClient = bufio.NewReader(conn)

	return tcpSRH
}

func (c *TCPServerRequestHandler) send(msg []byte) {
	encoder := gob.NewEncoder(c.inToClient)

	encoder.Encode(msg)

	c.inToClient.Flush()
}

func (c *TCPServerRequestHandler) receive() []byte {
	decoder := gob.NewDecoder(c.outToClient)

	var data []byte

	decoder.Decode(&data)

	return data
}
