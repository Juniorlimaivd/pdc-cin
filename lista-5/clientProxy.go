package main

type serviceInformation struct {
	host          string
	port          int
	objectId      int
	operationName string
	invocation    Invocation
	termination   Termination
	parameters    []string
	methodName    string
	//requestor     Requestor
}

func (s *serviceInformation) getHost() string {
	return s.host
}

func (s *serviceInformation) setHost(host string) {
	s.host = host
}

func (s *serviceInformation) getPort() int {
	return s.port
}

func (s *serviceInformation) setPort(port int) {
	s.port = port
}

func (s *serviceInformation) getObjectId() int {
	return s.objectId
}

func (s *serviceInformation) setObjectId(objectId int) {
	s.objectId = objectId
}

func (s *serviceInformation) getBalance(accountId string) string {
	s.methodName = "getBalance"
	s.parameters = append(s.parameters, accountId)

	s.invocation.setObjectId(s.getObjectId())
	s.invocation.setIpAdress(s.getHost())
	s.invocation.setPortNumber(s.getPort())
	s.invocation.setOperationName(s.methodName)
	s.invocation.setParameters(s.parameters)

	//termination = requestor.invoke(s.invocation)

	return s.termination.getResult()
}

func (s *serviceInformation) withdraw(accountId string, amount float32) string {
	s.methodName = "withdraw"
	s.parameters = append(s.parameters, accountId)
	s.parameters = append(s.parameters, amount)

	s.invocation.setObjectId(s.getObjectId())
	s.invocation.setIpAdress(s.getHost())
	s.invocation.setPortNumber(s.getPort())
	s.invocation.setOperationName(s.methodName)
	s.invocation.setParameters(s.parameters)

	//termination = requestor.invoke(s.invocation)

	return s.termination.getResult()
}

func (s *serviceInformation) deposit(accountId string, amount float32) string {
	s.methodName = "deposit"
	s.parameters = append(s.parameters, accountId)
	s.parameters = append(s.parameters, amount)

	s.invocation.setObjectId(s.getObjectId())
	s.invocation.setIpAdress(s.getHost())
	s.invocation.setPortNumber(s.getPort())
	s.invocation.setOperationName(s.methodName)
	s.invocation.setParameters(s.parameters)

	//termination = requestor.invoke(s.invocation)

	return s.termination.getResult()
}

func (s *serviceInformation) transfer(payerId string, payeeId string, amount float32) string {
	s.methodName = "transfer"
	s.parameters = append(s.parameters, payerId)
	s.parameters = append(s.parameters, payeeId)
	s.parameters = append(s.parameters, amount)

	s.invocation.setObjectId(s.getObjectId())
	s.invocation.setIpAdress(s.getHost())
	s.invocation.setPortNumber(s.getPort())
	s.invocation.setOperationName(s.methodName)
	s.invocation.setParameters(s.parameters)

	//termination = requestor.invoke(s.invocation)

	return s.termination.getResult()
}
