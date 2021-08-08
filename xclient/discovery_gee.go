package xclient

import "time"

const (
	defaultUpdateTimeout = time.Second * 10
)

type GeeRegistryDiscovery struct {
	*MultiServersDiscovery
	registry   string
	timeout    time.Duration
	lastUpdate time.Time
}