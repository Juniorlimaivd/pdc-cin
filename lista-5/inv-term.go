package main

type Invocation struct {
	objectId      int
	ipAdress      string
	portNumber    int
	operationName string
	parameters    []string
}

type Termination struct {
	result string //ver qual o tipo de result (em java é Object)
}

func (inv *Invocation) setObjectId(objectId int) {
	inv.objectId = objectId
}

func (inv *Invocation) getObjectId() int {
	return inv.objectId
}

func (inv *Invocation) setIpAdress(ipAdress string) {
	inv.ipAdress = ipAdress
}

func (inv *Invocation) getIpAdress() string {
	return inv.ipAdress
}

func (inv *Invocation) setPortNumber(portNumber int) {
	inv.portNumber = portNumber
}

func (inv *Invocation) getPortNumber() int {
	return inv.portNumber
}

func (inv *Invocation) setOperationName(operationName string) {
	inv.operationName = operationName
}

func (inv *Invocation) getOperationName() string {
	return inv.operationName
}

func (inv *Invocation) setParameters(parameters []int) {
	inv.parameters = parameters
}

func (inv *Invocation) getParameters() []string {
	return inv.parameters
}

func (term *Termination) setResult(result string) {
	term.result = result
}

func (term *Termination) getResult() string {
	return term.result
}
