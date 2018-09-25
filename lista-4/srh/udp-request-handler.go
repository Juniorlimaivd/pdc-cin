package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
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
	udpSRH.listener, err = net.ListenPacket("udp", ":"+port)
	if err != nil {
		fmt.Println(err.Error())
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
	_, c.clientAddr, err = c.listener.ReadFrom(buffer)

	if err != nil {
		fmt.Println(err.Error())
	}

	var data []byte

	reader := bytes.NewReader(buffer)
	dec := gob.NewDecoder(reader)
	err = dec.Decode(&data)

	if err != nil {
		fmt.Println(err.Error())
	}

	return data
}
