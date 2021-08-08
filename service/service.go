package service

import (
	"reflect"
)

type MethodType struct {
	method    reflect.Method
	ArgType   reflect.Type
	ReplyType reflect.Type
	numCalls  uint64
}

type Service struct {
	name   string
	typ    reflect.Type
	rcvr   reflect.Value
	method map[string]*MethodType
}
