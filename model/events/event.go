package events

type Event struct {
	stopPropagation bool
	data			interface{}
}

func (e *Event) StopPropagation() bool {
	return e.stopPropagation
}

func (e *Event) SetStopPropagation(stopPropagation bool) {
	e.stopPropagation = stopPropagation
}

func (e *Event) Data() interface{} {
	return e.data
}

func (e *Event) SetData(data interface{}) {
	e.data = data
}



