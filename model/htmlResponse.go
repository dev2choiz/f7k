package model

import (
	"fmt"
	"net/http"
)

type HtmlResponse struct {
	Response
	Content string
}

func (r *HtmlResponse) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/html")
	_, e := fmt.Fprintf(w, "%v", r.Content)
	if nil != e {
		return e
	}
	return nil
}
