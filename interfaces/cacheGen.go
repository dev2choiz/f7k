package interfaces

type CacheListener interface {
	Abort() bool
	SetAbort(bool) CacheListener

}
