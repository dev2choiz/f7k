package events

type ResponseEvent struct {
	KernelEvent
}

func (h *ResponseEvent) GetEventName() string {
	return "KernelResponseEvent"
}
