package serverrequesthandler

import (
	"errors"
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

func NewServerRequestHandler(handlerType string, port int) (*ServerRequestHandler, error) {
	srh := ServerRequestHandler{handlerType: handlerType, port: port}

	switch handlerType {
	case "udp":
		srh.handler = NewUDPServerRequestHandler(strconv.Itoa(port))
		break
	case "tcp":
		srh.handler = newTCPServerRequestHandler(strconv.Itoa(port))
		break
	case "middleware":
		srh.handler = NewAMQServerRequestHandler(strconv.Itoa(port))
		break
	default:
		return nil, errors.New("No handler of type \"" + handlerType + "\" found")
	}

	return &srh, nil

}

func (srh *ServerRequestHandler) send(data []byte) {
	srh.handler.send(data)
}

func (srh *ServerRequestHandler) receive() []byte {
	return srh.handler.receive()
}
