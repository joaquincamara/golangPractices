package api

import (
	"encoding/json"
	"net/http"

	"github.com/joaquincamara/golangPractices/pkg/controller"
	"github.com/joaquincamara/golangPractices/pkg/model"
)

type EventHandler struct {
	Handler  IHandler
	EventCtr controller.IEventController
}

var messages = Messages{
	Exist:             "Event already exist",
	InvalidRequest:    "Invalid request",
	EditionSucces:     "Edition succesful",
	DeleteSucces:      "Deletion succesful",
	CreationSucces:    "Creaction succesful",
	ElementsFound:     "Elements Found",
	SomethingGetWrong: "Something get wrong",
}

func NewEventHandler(eventCtr controller.IEventController) IHandler {
	return &EventHandler{EventCtr: eventCtr}
}

func (eh *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event model.Event
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&event); err != nil {
		WithErrorResponse(w, http.StatusBadRequest, messages.InvalidRequest)
		return
	}

	if err := eh.EventCtr.Add(event); err != nil {
		WithErrorResponse(w, http.StatusFound, messages.Exist)
		return
	}

	WithSuccesResponse(w, http.StatusCreated, messages.CreationSucces)
}

func (eh *EventHandler) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	list, err := eh.EventCtr.FindAll()
	if err != nil {
		WithErrorResponse(w, http.StatusBadRequest, messages.InvalidRequest)
	}

	response, _ := json.Marshal(list)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (eh *EventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	var event model.Event
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&event); err != nil {
		WithErrorResponse(w, http.StatusBadRequest, messages.InvalidRequest)
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			WithErrorResponse(w, http.StatusInternalServerError, messages.SomethingGetWrong)
		}
	}()

	if err := eh.EventCtr.Delete(event); err != nil {
		WithErrorResponse(w, http.StatusBadRequest, messages.InvalidRequest)
	}

	WithSuccesResponse(w, http.StatusOK, messages.DeleteSucces)
}

func (eh *EventHandler) EditEvent(w http.ResponseWriter, r *http.Request) {
	var event model.Event
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&event); err != nil {
		WithErrorResponse(w, http.StatusBadRequest, messages.InvalidRequest)
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			WithErrorResponse(w, http.StatusInternalServerError, messages.SomethingGetWrong)
		}
	}()

	if err := eh.EventCtr.Edit(event); err != nil {
		WithErrorResponse(w, http.StatusBadRequest, messages.InvalidRequest)
	}
	WithSuccesResponse(w, http.StatusOK, messages.EditionSucces)
}
