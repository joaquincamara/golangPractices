package api

import (
	"encoding/json"
	"net/http"
)

type Messages struct {
	Exist             string
	InvalidRequest    string
	EditionSucces     string
	DeleteSucces      string
	CreationSucces    string
	ElementsFound     string
	SomethingGetWrong string
}

func WithSuccesResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(message)
	w.WriteHeader(code)
	w.Write(response)
}

func WithErrorResponse(w http.ResponseWriter, code int, message string) {
	WithSuccesResponse(w, code, message)
}
