package interfaces

import "sync"

type Handler func(e Event)
type AsyncHandler func(listener AsyncListenerMetadata)

type EventDispatcher interface {
	Dispatch(eventName string, event Event)
	Listen(eventName string, handler Handler)

	// Async
	InitDispatcher(eventName, disId  string)
	DispatchAsync(eventName, dispName string, event AsyncEvent) (AsyncEventMetadata, error)
	ListenAsync(eventName, listenerName string) (chan AsyncEvent,AsyncListenerMetadata, error)

	StopDispatcher(eventName, dispatcherName string) error
	WaitUntilAsyncListeners(eventName string)
	CloseListener(eventName, listenerName string)
}

type AsyncListenerMetadata interface {
	Id() string
	Done()
	Payload() interface{}
	SetPayload(interface{})
	WaitGroup() *sync.WaitGroup
}

type AsyncEventMetadata interface {
	WaitGroup() *sync.WaitGroup
	Wait()
	Payload() interface{}
	SetPayload(interface{})
	HasListeners() bool
	SetHasListeners(bool)
	ListenersWaiter() chan bool
	AsyncListenerMetadata() map[string]AsyncListenerMetadata
	AllDispatchersEnd() bool
}