package distributed

import (
	"sync/atomic"
)

type Replication interface {
	Replicate(int) int
}

func NewRoundRobin() *RoundRobinSelector {
	return &RoundRobinSelector{0}
}

type RoundRobinSelector struct {
	index uint32
}

func (replica *RoundRobinSelector) Replicate(N int) int {
	index := atomic.AddUint32(&replica.index, 1)
	return int(index % uint32(N))
}
