package utils

import (
	"encoding/json"
	"net/http"
)

func WithSuccesResponse(w http.ResponseWriter, code int, message string) {
	response, _ := json.Marshal(message)
	w.WriteHeader(code)
	w.Write(response)
}

func WithErrorResponse(w http.ResponseWriter, code int, message string) {
	WithSuccesResponse(w, code, message)
}
