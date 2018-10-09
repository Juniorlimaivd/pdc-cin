package main

import (
	"fmt"
	"reflect"
)

type requestPkt struct {
	MethodName string
	Args       []interface{}
	ReturnType interface{}
}

type returnPkt struct {
	MethodName  string
	ReturnValue interface{}
	Err         error
}

func main() {
	crh := NewClientRequestHandler("localhost", 1234, "tcp")
	crh.connect()
	fmt.Scanln()
	accID := "AC1"
	accID2 := "AC2"
	var amount float32 = 100.0
	args := []interface{}{reflect.ValueOf(accID).Interface(), reflect.ValueOf(accID2).Interface(), reflect.ValueOf(amount).Interface()}

	var testfloat = "a"
	request := requestPkt{MethodName: "Transfer", Args: args, ReturnType: reflect.ValueOf(testfloat).Interface()}
	fmt.Println(request)
	marshaller := new(Marshaller)

	data := marshaller.marshall(request)
	crh.send(data)
	res := crh.receive()
	var resPkt returnPkt
	marshaller.unmarshall(res, &resPkt)
	v := reflect.ValueOf(resPkt.ReturnValue)
	fmt.Println(v)
	fmt.Scanln()
	fmt.Scanln()
	fmt.Scanln()
}
