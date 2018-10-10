package client

import (
	"../global"
)

// Requestor ...
type Requestor struct {
	crh        *ClientRequestHandler
	marshaller *global.Marshaller
}

func newRequestor(host string, port int) *Requestor {
	crh := newClientRequestHandler(host, port)
	crh.connect()
	return &Requestor{
		crh:        crh,
		marshaller: new(global.Marshaller),
	}
}

func (r *Requestor) invoke(request global.RequestPkt) *global.ReturnPkt {
	marshContent := r.marshaller.Marshall(request)
	r.crh.send(marshContent)

	marshRet := r.crh.receive()
	var resPkt global.ReturnPkt
	r.marshaller.Unmarshall(marshRet, &resPkt)

	return &resPkt
}
