package model

import (
	"net/http"
)

type Response struct {
	status uint16
}

func (r *Response) Send(w http.ResponseWriter) error {
	return nil
}

func (r *Response) Status() uint16 {
	return r.status
}

func (r *Response) SetStatus(status uint16) {
	r.status = status
}