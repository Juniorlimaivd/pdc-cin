package client

import (
	"reflect"

	"../global"
)

// Proxy ...
type Proxy struct {
	host      string
	port      int
	requestor *Requestor
}

func newProxy(host string, port int) *Proxy {
	return &Proxy{
		host:      host,
		port:      port,
		requestor: newRequestor(host, port),
	}
}

func (p *Proxy) getBalance(accountID string) float64 {
	reqPkt := global.NewRequestPkt("GetBalance", accountID)
	retPkt := p.requestor.invoke(reqPkt)
	return reflect.ValueOf(retPkt.ReturnValue).Float()
}

func (p *Proxy) withdraw(accountID string, amount float64) string {
	reqPkt := global.NewRequestPkt("Withdraw", accountID, amount)
	retPkt := p.requestor.invoke(reqPkt)
	return reflect.ValueOf(retPkt.ReturnValue).String()
}

func (p *Proxy) deposit(accountID string, amount float64) string {
	reqPkt := global.NewRequestPkt("Deposit", accountID, amount)
	retPkt := p.requestor.invoke(reqPkt)
	return reflect.ValueOf(retPkt.ReturnValue).String()
}

func (p *Proxy) transfer(payerID string, payeeID string, amount float64) string {
	reqPkt := global.NewRequestPkt("Transfer", payerID, payeeID, amount)
	retPkt := p.requestor.invoke(reqPkt)
	return reflect.ValueOf(retPkt.ReturnValue).String()
}
