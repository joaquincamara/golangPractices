package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joaquincamara/golangPractices/pkg/api"
	"github.com/joaquincamara/golangPractices/pkg/controller"
	"github.com/joaquincamara/golangPractices/pkg/database"
	"github.com/joaquincamara/golangPractices/pkg/model"
)

var DB = make(database.DB)

func main() {
	DB[1] = model.Event{
		Id:          1,
		Title:       "Buy food for the weekend",
		Description: "Remember the vegetables",
	}
	eventRepo := database.NewEventRepository(DB)
	eventCtr := controller.NewEventController(eventRepo)
	eventHandler := api.NewEventHandler(eventCtr)

	router := mux.NewRouter()
	router.HandleFunc("/", Home)
	router.HandleFunc("/api/v1/events", eventHandler.GetAllEvents).Methods("GET")
	router.HandleFunc("/api/v1/events", eventHandler.CreateEvent).Methods("POST")
	router.HandleFunc("/api/v1/events", eventHandler.EditEvent).Methods("PUT")
	router.HandleFunc("/api/v1/events", eventHandler.DeleteEvent).Methods("DELETE")

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
