package cacheGen

import (
	"github.com/dev2choiz/f7k"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/model/events"
	"sync"
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
	mux := &sync.Mutex{}
	for name, f := range WaitingForListen {
		go func(name string, f func(*events.CacheGenEvent)) {
			ch, l, err := d.ListenAsync(events.OnCacheGenEvent, name)
			if err != nil {
				panic(err)
			}
			<-ch
			ev := events.CacheGenEvent{}
			f(&ev)

			mux.Lock()
			pay := l.Payload().(*events.CacheGenEvent)
			mergeEvents(&ev, pay)
			l.SetPayload(&ev)
			l.Done()
			mux.Unlock()
			d.CloseListener(events.OnCacheGenEvent, name)
		}(name, f)
	}
}

func mergeEvents(a, b *events.CacheGenEvent) {
	a.ImportCachePackages = mergeSlices(a.ImportCachePackages, b.ImportCachePackages)
	a.GeneratedFiles = mergeSlices(a.GeneratedFiles, b.GeneratedFiles)
	a.PreAppLoadFunctions = mergeSlices(a.PreAppLoadFunctions, b.PreAppLoadFunctions)
	a.PostAppLoadFunctions = mergeSlices(a.PostAppLoadFunctions, b.PostAppLoadFunctions)
}

func mergeSlices(a, b []string) []string {
	news := make([]string, 0)
	for _, v := range append(a, b...) {
		exist := false
		for _, n := range news {
			if v == n {
				exist = true
				break
			}
		}
		if !exist {
			news = append(news, v)
		}
	}

	return news
}
