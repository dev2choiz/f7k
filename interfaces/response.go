package interfaces

import (
	"net/http"
)

type ResponseInterface interface {
	Send(w http.ResponseWriter) error
	Status() uint16
	SetStatus(status uint16)
}
