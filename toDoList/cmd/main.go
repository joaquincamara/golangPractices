package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Event struct {
	Id          int32
	Title       string
	Description string
}

var DB = make(map[int32]Event)

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
	fmt.Println("Welcome to the ToDo app instructions")
	fmt.Println("APIS:")
	fmt.Println("This is the model for a event the app")
	fmt.Println("Id: int32")
	fmt.Println("Title: string")
	fmt.Println("Description: string")

	fmt.Println("This is our endpoint and the methods that the url can manage")
	fmt.Println("GET, POST, PUT, DELETE")
	fmt.Println("/api/v1/events")
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	message := make(map[string]string)
	var event Event
	decoder := json.NewDecoder(r.Body)

	w.Header().Set("Content-Type", "application/json")

	if err := decoder.Decode(&event); err != nil {
		message["error"] = "Invalid request"
		response, _ := json.Marshal(message["error"])
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}
	defer r.Body.Close()

	if _, ok := DB[event.Id]; ok {
		message["Exist"] = "Event already exist"
		response, _ := json.Marshal(message)
		w.WriteHeader(http.StatusFound)
		w.Write(response)
	} else {
		DB[event.Id] = event
		response, _ := json.Marshal(DB[event.Id])
		w.WriteHeader(http.StatusCreated)
		w.Write(response)
	}

}

func EditEvent(w http.ResponseWriter, r *http.Request) {
	message := make(map[string]string)
	var event Event
	decoder := json.NewDecoder(r.Body)

	w.Header().Set("Content-Type", "application/json")

	if err := decoder.Decode(&event); err != nil {
		message["error"] = "Invalid request"
		response, _ := json.Marshal(message["error"])
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}
	defer r.Body.Close()

	if _, ok := DB[event.Id]; ok {
		message["EditionSucces"] = "Edition succesful"
		DB[event.Id] = event
		response, _ := json.Marshal(message["EditionSucces"])
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	} else {
		message["error"] = "Invalid request"
		response, _ := json.Marshal(message["error"])
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	message := make(map[string]string)
	var event Event
	decoder := json.NewDecoder(r.Body)

	w.Header().Set("Content-Type", "application/json")

	if err := decoder.Decode(&event); err != nil {
		message["error"] = "Invalid request"
		response, _ := json.Marshal(message["error"])
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}
	defer r.Body.Close()

	if _, ok := DB[event.Id]; ok {
		message["DeleteSucces"] = "Delete succesful"
		delete(DB, event.Id)
		response, _ := json.Marshal(message["DeleteSucces"])
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	} else {
		message["error"] = "Invalid request"
		response, _ := json.Marshal(message["error"])
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}
}

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(DB)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
