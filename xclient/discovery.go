package xclient

import (
	"math/rand"
	"sync"
)

type SelectMode int

const (
	RandomSelect     SelectMode = iota // select randomly
	RoundRobinSelect                   // select using Robbin algorithm
	WeightRoundRobin
	ConsistentHash
)

type Discovery interface {
	Refresh() error // refresh from remote registryΩ
	Update(servers []string) error
	Get(mode SelectMode) (string, error)
	GetAll() ([]string, error)
}

// MultiServersDiscovery is a discovery for multi servers without a registry center
// user provides the server addresses explicitly instead
type MultiServersDiscovery struct {
	r       *rand.Rand    // generate random number
	mu      sync.RWMutex  // protect following
	servers []string
	index   int  // record the selected position for robin algorithm
}