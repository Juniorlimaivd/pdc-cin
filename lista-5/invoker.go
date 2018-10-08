package middleware

import (
	"log"
	"reflect"

	"./serverrequesthandler"
)

type requestPkt struct {
	methodName string
	argTypes   []reflect.Value
	returnType reflect.Value
}

type returnPkt struct {
	methodName  string
	returnValue reflect.Value
	err         error
}

type methodInfo struct {
	method    reflect.Method
	argsType  []reflect.Type
	replyType reflect.Type
}

// Invoker handles the directing from ServerRequestHandler to the correct remote method
type Invoker struct {
	srh        *serverrequesthandler.ServerRequestHandler
	marshaller *Marshaller
	methods    map[string]*methodInfo
	object     interface{}
}

func newInvoker(object interface{}) *Invoker {
	srh, _ := serverrequesthandler.NewServerRequestHandler("tcp", 1234)
	inv := Invoker{srh: srh,
		marshaller: new(Marshaller),
		object:     object}

	inv.registerMethods()
	return &inv
}

func (i *Invoker) registerMethods() {
	methods := make(map[string]*methodInfo)

	objectType := reflect.TypeOf(i.object)

	for i := 0; i < objectType.NumMethod(); i++ {
		method := objectType.Method(i)
		methodName := method.Name
		methodType := method.Type

		argsType := []reflect.Type{}

		for j := 0; j < methodType.NumIn(); j++ {
			argsType = append(argsType, methodType.In(j))
		}

		if methodType.NumOut() != 1 {
			log.Printf("invoker.registerMethods: method %q has %d output parameters; needs exactly one\n", methodName, methodType.NumOut())
			continue
		}

		returnType := methodType.Out(0)

		methods[methodName] = &methodInfo{method: method, argsType: argsType, replyType: returnType}
	}
}

func (i *Invoker) handleOperation(request *requestPkt) returnPkt {
	methodInf := i.methods[request.methodName]

	if len(request.argTypes) != len(methodInf.argsType) {
		log.Printf("invoker.handleOperation: request has %d parameters; needs exactly %d\n", len(request.argTypes), len(methodInf.argsType))
	}

	in := []reflect.Value{methodInf.method.Func}
	for _, arg := range request.argTypes {
		in = append(in, arg)
	}
	in = append(in, request.returnType)
	methodInf.method.Func.Call(in)
	return returnPkt{methodName: request.methodName, returnValue: request.returnType, err: nil}
}

func (i *Invoker) invoke() {
	for {
		data := i.srh.Receive()
		request := new(requestPkt)
		i.marshaller.unmarshall(data, &request)
		go func() {
			returnPkt := i.handleOperation(request)
			pkt := i.marshaller.marshall(returnPkt)
			i.srh.Send(pkt)
		}()
	}
}
