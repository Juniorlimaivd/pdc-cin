package main

import (
	"net/rpc"
)

// RPCClientRequestHandler handles tcp connections
type RPCClientRequestHandler struct {
	host string
	port int

	client *rpc.Client
}

func newRPCClientRequestHandler(host string, port int) *RPCClientRequestHandler {
	return &RPCClientRequestHandler{
		host: host,
		port: port,
	}
}

func (c *RPCClientRequestHandler) connect() error {
	addr := c.host + ":" + string(c.port)
	var err error
	c.client, err = rpc.DialHTTP("tcp", addr)
	return err
}

func (c *RPCClientRequestHandler) send(data []byte) error {
	c.client.Call("Receiver.SendedByte", data, nil)
}

func (c *RPCClientRequestHandler) receive() []byte {
	var data []byte
	c.client.Call("Receiver.ReceiveByte", nil, &data)
	return data
}
