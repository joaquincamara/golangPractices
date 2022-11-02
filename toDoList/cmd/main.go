package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joaquincamara/golangPractices/pkg/api"
	"github.com/joaquincamara/golangPractices/pkg/database"
	"github.com/joaquincamara/golangPractices/pkg/event"
)

var DB database.DB

var messages = api.Messages{
	Exist:             "Event already exist",
	InvalidRequest:    "Invalid request",
	EditionSucces:     "Edition succesful",
	DeleteSucces:      "Deletion succesful",
	CreationSucces:    "Creaction succesful",
	ElementsFound:     "Elements Found",
	SomethingGetWrong: "Something get wrong",
}

func main() {

	eventRepo := database.NewEventRepository(DB)
	eventCtr := event.NewEventController(eventRepo)
	fmt.Println(eventCtr)

	router := mux.NewRouter()
	router.HandleFunc("/", Home)
	router.HandleFunc("/api/v1/events", GetAllEvents).Methods("GET")
	router.HandleFunc("/api/v1/events", CreateEvent).Methods("POST")
	router.HandleFunc("/api/v1/events", EditEvent).Methods("PUT")
	router.HandleFunc("/api/v1/events", DeleteEvent).Methods("DELETE")

	http.ListenAndServe(":8080", router)

}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println(`Welcome to the ToDo app instructions
	This is the model for a event the app

	Id: int32
	Title: string
	Title: string
	Description: string

	This is our endpoint and the methods that the url can manage
	GET, POST, PUT, DELETE
	/api/v1/events
	`)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event event.Event
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&event); err != nil {
		api.WithErrorResponse(w, http.StatusBadRequest, messages.InvalidRequest)
		return
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			api.WithErrorResponse(w, http.StatusInternalServerError, messages.SomethingGetWrong)
			return
		}
	}()

	if _, ok := DB[event.Id]; ok {
		api.WithErrorResponse(w, http.StatusFound, messages.Exist)
		return
	}

	DB[event.Id] = event
	api.WithSuccesResponse(w, http.StatusCreated, messages.CreationSucces)

}

func EditEvent(w http.ResponseWriter, r *http.Request) {
	var event event.Event
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&event); err != nil {
		api.WithErrorResponse(w, http.StatusBadRequest, messages.InvalidRequest)
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			api.WithErrorResponse(w, http.StatusInternalServerError, messages.SomethingGetWrong)
		}
	}()

	if _, ok := DB[event.Id]; ok {
		DB[event.Id] = event
		api.WithSuccesResponse(w, http.StatusOK, messages.EditionSucces)
	}

	api.WithErrorResponse(w, http.StatusBadRequest, messages.InvalidRequest)

}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	var event event.Event
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&event); err != nil {
		api.WithErrorResponse(w, http.StatusBadRequest, messages.InvalidRequest)
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			api.WithErrorResponse(w, http.StatusInternalServerError, messages.SomethingGetWrong)
		}
	}()

	if _, ok := DB[event.Id]; ok {

		delete(DB, event.Id)
		api.WithSuccesResponse(w, http.StatusOK, messages.DeleteSucces)
	}

	api.WithErrorResponse(w, http.StatusBadRequest, messages.InvalidRequest)
}

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(DB)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
