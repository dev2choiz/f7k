package eventDispatcher

import (
	"github.com/dev2choiz/f7k/interfaces"
	"sync"
)

/*
this variable is public so that it can be overridden by a struct that implements the interfaces.EventDispatcher interface
This has the effect of only using the new struct throughout the application
 */
var Dispatcher interfaces.EventDispatcher

type EventDispatcher struct {
	Listeners           map[string][]interfaces.Handler
	AsyncEventsMetadata map[string]interfaces.AsyncEventMetadata
	mux                 sync.Mutex
}

func Instance() interfaces.EventDispatcher {
	if nil == Dispatcher {
		d := &EventDispatcher{}
		d.Listeners = make(map[string][]interfaces.Handler)
		d.AsyncEventsMetadata = map[string]interfaces.AsyncEventMetadata{}
		Dispatcher = d
	}

	return Dispatcher
}

func (ed *EventDispatcher) Listen(eventName string, handler interfaces.Handler) {
	if _, ok := ed.Listeners[eventName]; !ok {
		ed.Listeners[eventName] = []interfaces.Handler{}
	}
	ed.Listeners[eventName] = append(ed.Listeners[eventName], handler)
}

func (ed *EventDispatcher) Dispatch(eventName string, event interfaces.Event) {
	if _, ok := ed.Listeners[eventName]; ok {
		for _, handler := range ed.Listeners[eventName] {
			handler(event)
			if event.StopPropagation() {
				break
			}
		}
	}
}
