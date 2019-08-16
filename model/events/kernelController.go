package events

type ControllerEvent struct {
	KernelEvent
}

func (h *ControllerEvent) GetEventName() string {
	return "KernelControllerEvent"
}
