package client

import (
	"GeeRPC/codec"
	"errors"
	"io"
	"net"
	"sync"
)

var (
	_           io.Closer = (*Client)(nil)
	ErrShutdown           = errors.New("connection is shut down")
)

const (
	defaultRPCPath   = "/_geeprc_"
	connected        = 200
	defaultDebugPath = "/debug/geerpc"
)

// Call represents an active RPC.
type Call struct {
	Seq           uint64
	ServiceMethod string      // format "<service>.<method>"
	Args          interface{} // arguments to the function
	Reply         interface{} // reply from the function
	Error         error       // if error occurs, it will be set
	Done          chan *Call  // Strobes when call is complete.
}

// Client represents an RPC Client.
// There may be multiple outstanding Calls associated
// with a single Client, and a Client may be used by
// multiple goroutines simultaneously.
type Client struct {
	cc       codec.Codec
	opt      *codec.Option
	sending  sync.Mutex
	header   codec.Header
	mu       sync.Mutex
	seq      uint64
	pending  map[uint64]*Call
	closing  bool
	shutdown bool
}

type clientResult struct {
	client *Client
	err    error
}

type newClientFunc func(conn net.Conn, opt *codec.Option) (client *Client, err error)
