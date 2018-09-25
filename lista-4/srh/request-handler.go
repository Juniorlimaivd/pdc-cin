package main

type ServerRequestHandler interface {
	send(msg []byte)
	receive() []byte
}
