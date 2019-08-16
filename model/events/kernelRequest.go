package events

type RequestEvent struct {
	KernelEvent
}

func (h *RequestEvent) GetEventName() string {
	return "KernelRequestEvent"
}

