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
	outToClient *bufio.Reader
	inToClient  *bufio.Writer
}

func newTCPServerRequestHandler(port string) *TCPServerRequestHandler {
	tcpSRH := new(TCPServerRequestHandler)
	tcpSRH.listener, _ = net.Listen("tcp", port)

	log.Println("Listen on", tcpSRH.listener.Addr().String())
	conn, _ := tcpSRH.listener.Accept()
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
