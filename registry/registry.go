package registry

import (
"sync"
"time"
)

//                 |registry|
//           pull /push      \register
//          |Client|   ->   |Server|
//                    call

//服务端启动后，向注册中心发送注册消息，注册中心得知该服务已经启动，处于可用状态。一般来说，服务端还需要定期向注册中心发送心跳，证明自己还活着。
//客户端向注册中心询问，当前哪天服务是可用的，注册中心将可用的服务列表返回客户端。
//客户端根据注册中心得到的服务列表，选择其中一个发起调用。

const (
	defaultPath    = "/_geerpc_/registry"
	defaultTimeout = time.Minute * 5
)

// GeeRegistry is a simple register center, provide following functions.
// add a server and receive heartbeat to keep it alive.
// returns all alive servers and delete dead servers sync simultaneously.
type GeeRegistry struct {
	timeout time.Duration
	mu      sync.Mutex
	servers map[string]*ServerItem
}

type ServerItem struct {
	Addr  string
	start time.Time
}

