package main

func main() {
	handlerType := "middleware"
	port := "1234"
	switch handlerType {
	case "tcp":
		NewTCPServerRequestHandler(port)
		break
	case "udp":
		NewUDPServerRequestHandler(port)
		break
	case "middleware":
		NewRPCServerRequestHandler(port)
		break
	}
}
