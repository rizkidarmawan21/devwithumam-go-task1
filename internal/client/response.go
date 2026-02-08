package client

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(status int, message string, data interface{}) *Response {
	return &Response{Status: int(status), Message: message, Data: data}
}

// WriteJSON sets Content-Type: application/json and writes the standard API response.
// Use this so clients (e.g. Postman) always interpret the body as JSON.
func WriteJSON(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(NewResponse(status, message, data))
}
