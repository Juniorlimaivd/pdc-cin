package main

// ClientRequestHandler is a facade for connection handlers
type ClientRequestHandler struct {
	host               string
	port               int
	sentMessageSize    int
	receiveMessageSize int
	handlerType        string

	tcpHandler *TCPClientRequestHandler
	udpHandler *UDPClientRequestHandler
	midHandler *RPCClientRequestHandler
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
	switch crh.handlerType {
	case "tcp":
		crh.tcpHandler = newTCPClientRequestHandler(crh.host, crh.port)
		crh.tcpHandler.connect()
	case "udp":
		crh.udpHandler = newUDPClientRequestHandler(crh.host, crh.port)
		crh.udpHandler.connect()
		break
	case "middleware":
		crh.midHandler = newRPCClientRequestHandler(crh.host, crh.port)
		crh.midHandler.connect()
		break
	}
	return nil
}

func (crh *ClientRequestHandler) send(data []byte) {
	switch crh.handlerType {
	case "tcp":
		crh.tcpHandler.send(data)
		break
	case "udp":
		crh.udpHandler.send(data)
		break
	case "middleware":
		crh.midHandler.send(data)
		break
	}
}

func (crh *ClientRequestHandler) receive() []byte {
	switch crh.handlerType {
	case "tcp":
		return crh.tcpHandler.receive()
	case "udp":
		return crh.udpHandler.receive()
	case "middleware":
		return crh.midHandler.receive()
	}
	return nil
}
