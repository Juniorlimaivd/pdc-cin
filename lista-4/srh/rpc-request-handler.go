package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type RPCServerRequestHandler struct {
	port int
}

func NewRPCServerRequestHandler(port string) *RPCServerRequestHandler {
	rpcServerRequestHandler := new(RPCServerRequestHandler)
	rpc.Register(rpcServerRequestHandler)
	addr := ":" + port
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	return rpcServerRequestHandler
}

func (c *RPCServerRequestHandler) send(arg []byte, reply *string) {

}

func (c *RPCServerRequestHandler) receive() {

}
