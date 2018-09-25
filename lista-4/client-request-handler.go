package main

import (
	"net"

	"github.com/pkg/errors"
)

type ClientRequestHandler struct {
	host               string
	port               int
	sentMessageSize    int
	receiveMessageSize int
	handlerType        string
	tcpConnection      net.Conn
}

func newClientRequestHandler(host string, port int, handlerType string) *ClientRequestHandler {
	return &ClientRequestHandler{
		host:               host,
		port:               port,
		sentMessageSize:    1024,
		receiveMessageSize: 1024,
		handlerType:        handlerType,
	}
}

func (crh *ClientRequestHandler) connect() error {
	addr := crh.host + ":" + string(crh.port)
	switch crh.handlerType {
	case "tcp":
		tcpConnection, err := net.Dial("tcp", addr)
		if err != nil {
			return errors.Wrap(err, "Dialing "+addr+" failed")
		}
		break
	case "udp":
		break
	case "middleware":
		break
	}
}
