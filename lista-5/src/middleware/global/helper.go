package global

import (
	"reflect"
)

// RequestPkt ...
type RequestPkt struct {
	MethodName string
	Args       []interface{}
}

// NewRequestPkt ...
func NewRequestPkt(methodName string, args ...interface{}) RequestPkt {
	var _args []interface{}
	for _, arg := range args {
		_arg := reflect.ValueOf(arg).Interface()
		_args = append(_args, _arg)
	}
	return RequestPkt{
		MethodName: methodName,
		Args:       _args,
	}
}

// ReturnPkt ...
type ReturnPkt struct {
	MethodName  string
	ReturnValue interface{}
	Err         error
}
