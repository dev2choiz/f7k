package interfaces

type Handler func(e Event)
type AsyncHandler func(ch chan Event)

type EventDispatcher interface {
	Dispatch(eventName string, event Event)
	Listen(eventName string, handler Handler)

	// Async
	DispatchAsync(eventName string, event Event) error
	DispatchAsyncFunc(eventName string, handler AsyncHandler) error
	ListenAsync(eventName, listenerId string) (Event, error)
}
