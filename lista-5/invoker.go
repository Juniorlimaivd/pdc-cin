package middleware

import (
	"./serverrequesthandler"
)

type RequestPkt struct {
	serviceName string
	methodName string
	argTypes {}interface
	args {}interface
	returnTypes {}interface
}

type ReturnPkt struct {
	serviceName string
	methodName string 
	returnData {}interface
	err error 
}

type Invoker struct {
	srh serverrequeshandler.ServerRequestHandler
	marshaller *Marshaller
}

func newInvoker() *Invoker {
	return Invoker{srh: serverrequesthandler.NewServerRequestHandler("tcp", 1234),
						marshaller: new(Marshaller) }
}

func(i* Invoker) handleOperation(request *RequestPkt) ReturnPkt{

} 

func (i* Invoker) invoke(){
	for {
		data := i.srh.receive()
		request := new(RequestPkt)
		i.marshaller.unmarshall(data,&request)
		go func() {
			returnPkt := i.handleOperation(request)
			pkt := i.marshaller.marshall(returnPkt)
			i.srh.send(pkt)
		}
	}
}

