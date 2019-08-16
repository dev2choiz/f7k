package interfaces

import (
	"net/http"
)

type Kernel interface {
	Handle(w http.ResponseWriter, r *http.Request)
	RetrieveRoute() (Route, KernelEvent, *error)
	GetController(r Route) (ControllerInterface, KernelEvent, *error)
	CallController(ctrl ControllerInterface, r Route) (KernelEvent, *error)
	Finish(r ResponseInterface)
	ListenRequests()
}
