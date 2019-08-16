package model

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	Response
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func (hr *JsonResponse) init() {
	hr.Success = true
	hr.Message = ""
	hr.Data = nil
}

func (hr *JsonResponse) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(hr)
}
