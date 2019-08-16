package events

const OnCacheGenEvent = "on_cache_generate_event"

type CacheGenEvent struct {
	Event
	PreAppLoadFunctions  []string
	PostAppLoadFunctions []string
	ImportCachePackages  []string
	GeneratedFiles       []string
}

func (h *CacheGenEvent) GetEventName() string {
	return "CacheGenEvent"
}
