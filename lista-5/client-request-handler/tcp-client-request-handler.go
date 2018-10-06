package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"strconv"

	"github.com/pkg/errors"
)

// TCPClientRequestHandler handles tcp connections
type TCPClientRequestHandler struct {
	host               string
	port               int
	sentMessageSize    int
	receiveMessageSize int
	conn               net.Conn
	rw                 *bufio.ReadWriter
}

func newTCPClientRequestHandler(host string, port int) *TCPClientRequestHandler {
	return &TCPClientRequestHandler{
		host: host,
		port: port,
	}
}

func (c *TCPClientRequestHandler) connect() error {
	addr := c.host + ":" + strconv.Itoa(c.port)
	var err error
	c.conn, err = net.Dial("tcp", addr)

	if err != nil {
		return errors.Wrap(err, "Dialing "+addr+" failed")
	}

	fmt.Println("connected!")

	c.rw = bufio.NewReadWriter(bufio.NewReader(c.conn), bufio.NewWriter(c.conn))
	return nil
}

func (c *TCPClientRequestHandler) send(data []byte) error {

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

func (c *TCPClientRequestHandler) receive() []byte {
	var data []byte
	decoder := gob.NewDecoder(c.rw)
	decoder.Decode(&data)
	return data
}
