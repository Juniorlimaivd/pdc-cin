package main

import (
	"bufio"
	"log"
	"net"
)

type UDPServerRequestHandler struct {
	port        int
	listener    net.Listener
	rw          *bufio.ReadWriter
	outToClient *bufio.Reader
	inToClient  *bufio.Writer
}

func NewUDPServerRequestHandler(port string) *UDPServerRequestHandler {
	udpSRH := new(UDPServerRequestHandler)
	var err error
	udpSRH.listener, err = net.Listen("udp", port)
	if err != nil {
		// return errors.Wrapf(err, "Unable to listen on port %s\n", port)
	}
	log.Println("Listen on", udpSRH.listener.Addr().String())
	conn, err := udpSRH.listener.Accept()
	log.Println("Accept a connection request from", conn.RemoteAddr())

	udpSRH.inToClient = bufio.NewWriter(conn)
	udpSRH.outToClient = bufio.NewReader(conn)

	return udpSRH
}

func (c *UDPServerRequestHandler) send(arg []byte, reply *string) {

}

func (c *UDPServerRequestHandler) receive() {

}
