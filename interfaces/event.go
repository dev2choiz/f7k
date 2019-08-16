package interfaces

type Event interface {
	StopPropagation() bool
	SetStopPropagation(stopPropagation bool)
	Data() interface{}
	SetData(data interface{})
}
