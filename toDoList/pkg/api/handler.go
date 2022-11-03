package api

import "net/http"

type IHandler interface {
	CreateEvent(w http.ResponseWriter, r *http.Request)
	GetAllEvents(w http.ResponseWriter, r *http.Request)
	DeleteEvent(w http.ResponseWriter, r *http.Request)
	EditEvent(w http.ResponseWriter, r *http.Request)
}
