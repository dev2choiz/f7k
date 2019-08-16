package interfaces

import "net/http"

type KernelEvent interface {
	GetEventName() string
	GetRequest () *http.Request
	SetRequest (r *http.Request)
	GetResponseWriter () *http.ResponseWriter
	SetResponseWriter (r *http.ResponseWriter)
	GetController () ControllerInterface
	SetController (c ControllerInterface)
	GetResponse () ResponseInterface
	SetResponse (r ResponseInterface)
}
