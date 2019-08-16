package overrider

import (
	"github.com/dev2choiz/f7k"
	"github.com/dev2choiz/f7k/eventDispatcher"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/kernel"
	"github.com/dev2choiz/f7k/router"
)

func LoadEventDispatcherInstance() interfaces.EventDispatcher {
	if nil != f7k.Dispatcher {
		return f7k.Dispatcher
	}
	return eventDispatcher.Instance()
}

func LoadRouterInstance() interfaces.Router {
	if nil != f7k.Router {
		return f7k.Router
	}
	return router.Instance()
}

func LoadKernelInstance() interfaces.Kernel {
	if nil != f7k.Kernel {
		return f7k.Kernel
	}
	return kernel.Instance()
}

