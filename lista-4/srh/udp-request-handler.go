package main

import (
	"bufio"
	"log"
	"net"
)

type UDPServerRequestHandler struct {
	port       int
	listener   net.PacketConn
	rw         *bufio.ReadWriter
	clientAddr net.Addr
}

func NewUDPServerRequestHandler(port string) *UDPServerRequestHandler {
	udpSRH := new(UDPServerRequestHandler)
	var err error
	udpSRH.listener, err = net.ListenPacket("udp", port)
	if err != nil {
		// return errors.Wrapf(err, "Unable to listen on port %s\n", port)
	}
	log.Println("Listen on", udpSRH.listener.LocalAddr().String())

	return udpSRH
}

func (c *UDPServerRequestHandler) send(msg []byte) {
	c.listener.WriteTo(msg, c.clientAddr)
}

func (c *UDPServerRequestHandler) receive() []byte {
	buffer := make([]byte, 1024)
	var err error
	var n int
	n, c.clientAddr, err = c.listener.ReadFrom(buffer)
	return buffer[:n]
}
