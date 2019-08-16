package cacheGen

import (
	"github.com/dev2choiz/f7k"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/model/events"
)

type CacheListener struct {
	abort bool
}

func NewCacheListener() *CacheListener {
	return &CacheListener{}
}

func (cl *CacheListener) Abort() bool {
	return cl.abort
}

func (cl *CacheListener) SetAbort(a bool) interfaces.CacheListener {
	cl.abort = a

	return cl
}

// plug all listeners wanted to listen the cache event for write their owns cache files
func LoadCacheGenerators() {
	d := f7k.Dispatcher
	for _, f := range WaitingForListen {
		d.Listen(events.OnCacheGenEvent, f)
	}
}
