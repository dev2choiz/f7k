package events

type ConfigEvent struct {
	KernelEvent
}

func (c *ConfigEvent) GetEventName() string {
	return "ConfigEvent"
}
