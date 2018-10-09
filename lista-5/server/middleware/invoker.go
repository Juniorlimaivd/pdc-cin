package middleware

import (
	"fmt"
	"log"
	"reflect"

	"./serverrequesthandler"
)

type requestPkt struct {
	MethodName string
	Args       []interface{}
	ReturnType interface{}
}

type request struct {
	MethodName string
	Args       []reflect.Value
	ReturnType reflect.Type
}

type returnPkt struct {
	MethodName  string
	ReturnValue interface{}
	Err         error
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

// NewInvoker creates a new invoker
func NewInvoker(object interface{}) *Invoker {

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
		fmt.Println("methodName : ", methodName)
		fmt.Println("ArgsType : ", argsType)
		fmt.Println("ReturnType : ", returnType)
		methods[methodName] = &methodInfo{method: method, argsType: argsType, replyType: returnType}
	}
	i.methods = methods
}

func (i *Invoker) handleOperation(request *request) returnPkt {

	methodInf := i.methods[request.MethodName]

	if len(request.Args)+1 != len(methodInf.argsType) {
		log.Printf("invoker.handleOperation: request has %d parameters; needs exactly %d\n", len(request.Args), len(methodInf.argsType))
	}

	if request.ReturnType != methodInf.replyType {
		log.Printf("invoker.handleOperation: request asks for %+v type; needs %+v", request.ReturnType, methodInf.replyType)
	}

	in := []reflect.Value{reflect.ValueOf(i.object)}
	for _, arg := range request.Args {
		in = append(in, arg)
	}

	resultv := methodInf.method.Func.Call(in)
	return returnPkt{MethodName: request.MethodName, ReturnValue: resultv[0].Interface(), Err: nil}
}

func (i *Invoker) handleRequestPkt(requestPkt *requestPkt) *request {
	args := []reflect.Value{}
	for _, i := range requestPkt.Args {
		args = append(args, reflect.ValueOf(i))
	}

	res := reflect.TypeOf(requestPkt.ReturnType)

	return &request{MethodName: requestPkt.MethodName, Args: args, ReturnType: res}
}

// Invoke invokes the invoker
func (i *Invoker) Invoke() {
	for {
		data := i.srh.Receive()
		request := new(requestPkt)
		i.marshaller.unmarshall(data, &request)
		req := i.handleRequestPkt(request)
		fmt.Println(request)
		returnPkt := i.handleOperation(req)
		pkt := i.marshaller.marshall(returnPkt)
		i.srh.Send(pkt)
	}
}
