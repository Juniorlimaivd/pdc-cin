package main

import (
	"bufio"
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

func (c *TCPServerRequestHandler) send(arg []byte, reply *string) {

}

func (c *TCPServerRequestHandler) receive() {

}
