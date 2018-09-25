package main

import (
	"net/rpc"
	"strconv"
)

// RPCClientRequestHandler handles tcp connections
type RPCClientRequestHandler struct {
	host   string
	port   int
	id     string
	client *rpc.Client
}

func newRPCClientRequestHandler(host string, port int) *RPCClientRequestHandler {
	return &RPCClientRequestHandler{
		host: host,
		port: port,
	}
}

func (c *RPCClientRequestHandler) connect() error {
	addr := c.host + ":" + strconv.Itoa(c.port)
	var err error
	c.client, err = rpc.DialHTTP("tcp", addr)
	c.client.Call("Receiver.getID", nil, &c.id)
	return err
}

func (c *RPCClientRequestHandler) send(data []byte) error {
	c.client.Call("Receiver.SendedByte", data, nil)
	return nil
}

func (c *RPCClientRequestHandler) receive() []byte {
	var data []byte
	c.client.Call("Receiver.ReceiveByte", c.id, &data)
	return data
}
