package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"strconv"

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
	addr := c.host + ":" + strconv.Itoa(c.port)
	var err error

	c.conn, err = net.Dial("udp", addr)

	if err != nil {
		fmt.Println(err.Error())
		return errors.Wrap(err, "Dialing "+addr+" failed")
	}

	fmt.Println("connected!")
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
