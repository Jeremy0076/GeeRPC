package server

import (
	"GeeRPC/codec"
	"GeeRPC/service"
	"reflect"
	"sync"
)

const (
	connected        = "200 Connected to Gee RPC"
	defaultRPCPath   = "/_geerpc_"
	defaultDebugPath = "/debug/geerpc"
)

var DefaultServer = NewServer()

type Server struct {
	serviceMap sync.Map
}

// request stores all information of a call
type request struct {
	h            *codec.Header // header of request
	argv, replyv reflect.Value // argv and replyv of request
	mtype        *service.MethodType
	svc          *service.Service
}
