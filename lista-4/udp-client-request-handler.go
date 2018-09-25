package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"

	"github.com/pkg/errors"
)

// UDPClientRequestHandler handles tcp connections
type UDPClientRequestHandler struct {
	host string
	port int

	conn net.Conn
	rw   *bufio.ReadWriter
}

func newUDPClientRequestHandler(host string, port int) *UDPClientRequestHandler {
	return &UDPClientRequestHandler{
		host: host,
		port: port,
	}
}

func (c *UDPClientRequestHandler) connect() error {
	addr := c.host + ":" + string(c.port)
	var err error

	c.conn, err = net.Dial("udp", addr)

	if err != nil {
		return errors.Wrap(err, "Dialing "+addr+" failed")
	}
	c.rw = bufio.NewReadWriter(bufio.NewReader(c.conn), bufio.NewWriter(c.conn))
	return nil
}

func (c *UDPClientRequestHandler) send(data []byte) error {
	enc := gob.NewEncoder(c.rw)
	err := enc.Encode(data)
	if err != nil {
		fmt.Println("Error encoding", err)
		return err
	}
	err = c.rw.Flush()
	if err != nil {
		fmt.Println("Flush failed")
		return err
	}
	return nil
}

func (c *UDPClientRequestHandler) receive() []byte {
	var data []byte
	decoder := gob.NewDecoder(c.rw)
	decoder.Decode(&data)
	return data
}
