package main

import (
	"fmt"
)

func main() {
	e := NewEndpoint()
	e.Start()
	fmt.Scanln()
}
