package xclient

import (
	"GeeRPC/client"
	"GeeRPC/codec"
	"sync"
)

type XClient struct {
	d       Discovery
	mode    SelectMode
	opt     *codec.Option
	mu      sync.Mutex
	clients map[string]*client.Client
}