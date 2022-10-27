package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joaquincamara/golangPractices/utils"
)

type Event struct {
	Id          int32
	Title       string
	Description string
}

type Messages struct {
	Exist             string
	InvalidRequest    string
	EditionSucces     string
	DeleteSucces      string
	CreationSucces    string
	ElementsFound     string
	SomethingGetWrong string
}

var DB = make(map[int32]Event)
var messages = Messages{
	Exist:             "Event already exist",
	InvalidRequest:    "Invalid request",
	EditionSucces:     "Edition succesful",
	DeleteSucces:      "Deletion succesful",
	CreationSucces:    "Creaction succesful",
	ElementsFound:     "Elements Found",
	SomethingGetWrong: "Something get wrong",
}

func main() {
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
	var event Event
	decoder := json.NewDecoder(r.Body)

	w.Header().Set("Content-Type", "application/json")

	if err := decoder.Decode(&event); err != nil {
		utils.WithErrorResponse(w, http.StatusBadRequest, messages.InvalidRequest)
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			utils.WithErrorResponse(w, http.StatusInternalServerError, messages.SomethingGetWrong)
		}
	}()

	if _, ok := DB[event.Id]; ok {
		utils.WithErrorResponse(w, http.StatusFound, messages.Exist)
	}

	DB[event.Id] = event
	utils.WithSuccesResponse(w, http.StatusCreated, messages.CreationSucces)

}

func EditEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	decoder := json.NewDecoder(r.Body)

	w.Header().Set("Content-Type", "application/json")

	if err := decoder.Decode(&event); err != nil {
		utils.WithErrorResponse(w, http.StatusBadRequest, messages.InvalidRequest)
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			utils.WithErrorResponse(w, http.StatusInternalServerError, messages.SomethingGetWrong)
		}
	}()

	if _, ok := DB[event.Id]; ok {
		DB[event.Id] = event
		utils.WithSuccesResponse(w, http.StatusOK, messages.EditionSucces)
	}

	utils.WithErrorResponse(w, http.StatusBadRequest, messages.InvalidRequest)

}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	decoder := json.NewDecoder(r.Body)

	w.Header().Set("Content-Type", "application/json")

	if err := decoder.Decode(&event); err != nil {
		utils.WithErrorResponse(w, http.StatusBadRequest, messages.InvalidRequest)
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			utils.WithErrorResponse(w, http.StatusInternalServerError, messages.SomethingGetWrong)
		}
	}()

	if _, ok := DB[event.Id]; ok {

		delete(DB, event.Id)
		utils.WithSuccesResponse(w, http.StatusOK, messages.DeleteSucces)
	}

	utils.WithErrorResponse(w, http.StatusBadRequest, messages.InvalidRequest)
}

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(DB)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
