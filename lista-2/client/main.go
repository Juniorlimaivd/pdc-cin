package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
	}
}
