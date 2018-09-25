package main

import (
	"strconv"
)

// ServerRequestHandlerInterface interfaces server methods
type ServerRequestHandlerInterface interface {
	send(msg []byte)
	receive() []byte
}

type ServerRequestHandler struct {
	handler     ServerRequestHandlerInterface
	handlerType string
	port        int
}

func newServerRequestHandler(handlerType string, port int) *ServerRequestHandler {
	srh := ServerRequestHandler{handlerType: handlerType, port: port}

	switch handlerType {
	case "udp":
		srh.handler = NewUDPServerRequestHandler(strconv.Itoa(port))
		break
	case "tcp":
		srh.handler = newTCPServerRequestHandler(strconv.Itoa(port))
		break
	case "middleware":
		srh.handler = NewRPCServerRequestHandler(strconv.Itoa(port))
		break
	}

	return &srh

}

func (srh *ServerRequestHandler) send(data []byte) {
	srh.handler.send(data)
}

func (srh *ServerRequestHandler) receive() []byte {
	srh.handler.receive()
}
