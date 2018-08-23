package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

type Endpoint struct {
	ln   net.Listener
	conn net.Conn
	w    sync.WaitGroup
	ch   chan string
}

func NewEndpoint() *Endpoint {
	ep := &Endpoint{}
	return ep
}

func (ep *Endpoint) Start() {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	ep.ln = ln
	ep.ch = make(chan string)
	ep.w.Add(1)

	go func() {
		defer ep.w.Done()
		for {
			conn, err := ep.ln.Accept()
			if err != nil {
				fmt.Println("closed listener")
				break
			}
			fmt.Println("New client", ln.Addr())
			for {
				msg, err := bufio.NewReader(conn).ReadString('\n')
				fmt.Print("New Message:", msg)
				if err != nil {
					fmt.Println("closed socket")
					conn.Close()
					break
				}
				ep.ch <- msg
			}
		}
	}()

}

func (ep *Endpoint) Stop() {
	fmt.Println("stop endpoint")
	ep.ln.Close()
	if ep.conn != nil {
		ep.conn.Close()
	}
	ep.w.Wait()
}
