package eventDispatcher

import (
	"github.com/dev2choiz/f7k/interfaces"
)

/**
	Dispatch a event struct through a event channel.
	You can dispatch several events struct for a same event.
	To tell the event is finished, give nil instead of event struct as second argument.
	Nothing will be done if there isn't listener registered for the event.
 */
func (ed *EventDispatcher) DispatchAsync(eventName string, event interfaces.Event) error {
	ed.initEvent(eventName)
	for listenerName := range ed.AsyncEvents[eventName] {
		if nil == event {
			if nil != ed.AsyncEvents[eventName][listenerName] {
				close(ed.AsyncEvents[eventName][listenerName])
			}
			continue
		}
		ed.initListener(eventName, listenerName)
		ed.AsyncEvents[eventName][listenerName] <- event
	}
	return nil
}

/**
Dispatch an event struct through an event channel in your AsyncHandler function.
You can dispatch several events struct for a same event.
To tell the event is finished, send nil instead of event struct in the channel.
Nothing will be done if there isn't listener registered for the event.
*/
func (ed *EventDispatcher) DispatchAsyncFunc(eventName string, handler interfaces.AsyncHandler) error {
	for _, ch := range ed.AsyncEvents[eventName] {
		go func(handler interfaces.AsyncHandler, ch chan interfaces.Event) {
			handler(ch)
			if nil != ch {
				close(ch)
			}
		}(handler, ch)
	}
	return nil
}

/**
This function return an event received from the service dispatching.
You need to call it several times until you receive a nil value instead of event.
Do not listen a event if you are not sure it will be dispatched, because it's blocking.
*/
func (ed *EventDispatcher) ListenAsync(eventName, listenerName string) (interfaces.Event, error) {
	ed.initListener(eventName, listenerName)
	ch, _ := ed.AsyncEvents[eventName][listenerName]

	return <-ch, nil
}

func (ed *EventDispatcher) initEvent(eventName string) {
	if _, ok := ed.AsyncEvents[eventName]; !ok {
		ed.AsyncEvents[eventName] = map[string]chan interfaces.Event{}
	}
}

func (ed *EventDispatcher) initListener(eventName, listenerName string) {
	ed.initEvent(eventName)
	if _, ok := ed.AsyncEvents[eventName][listenerName]; !ok {
		ed.AsyncEvents[eventName][listenerName] = make(chan interfaces.Event)
	}
}
