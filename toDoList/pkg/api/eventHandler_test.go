package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/joaquincamara/golangPractices/pkg/controller"
	"github.com/joaquincamara/golangPractices/pkg/database"
	"github.com/joaquincamara/golangPractices/pkg/model"
	"github.com/stretchr/testify/assert"
)

var mockDB = make(database.DB)

// Mock json request
var eventJson = `{
	"Id": 1,
	"Title": "Buy food for the weekend",
	"Description": "Remember the vegetables"
	}`

func TestCreateEvent(t *testing.T) {
	mockEventRepo := database.NewEventRepository(mockDB)
	mockEventCtr := controller.NewEventController(mockEventRepo)
	mockEventHandler := NewEventHandler(mockEventCtr)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/events", strings.NewReader(eventJson))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockEventHandler.CreateEvent(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateEventExistingEvent(t *testing.T) {
	mockDB[1] = model.Event{
		Id:          1,
		Title:       "Buy food for the weekend",
		Description: "Remember the vegetables",
	}
	mockEventRepo := database.NewEventRepository(mockDB)
	mockEventCtr := controller.NewEventController(mockEventRepo)
	mockEventHandler := NewEventHandler(mockEventCtr)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/events", strings.NewReader(eventJson))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockEventHandler.CreateEvent(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
}

func TestEditEvent(t *testing.T) {
	mockDB[1] = model.Event{
		Id:          1,
		Title:       "Buy food for the weekend",
		Description: "Remember the vegetables",
	}

	var eventJsonEdit = `{
		"Id": 1,
		"Title": "Buy dog food for the weekend",
		"Description": "Remember the vegetables"
		}`
	mockEventRepo := database.NewEventRepository(mockDB)
	mockEventCtr := controller.NewEventController(mockEventRepo)
	mockEventHandler := NewEventHandler(mockEventCtr)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/events", strings.NewReader(eventJsonEdit))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockEventHandler.EditEvent(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestEditEventEventNotExist(t *testing.T) {
	var eventJsonEdit = `{
		"Id": 1,
		"Title": "Buy dog food for the weekend",
		"Description": "Remember the vegetables"
		}`
	mockEventRepo := database.NewEventRepository(mockDB)
	mockEventCtr := controller.NewEventController(mockEventRepo)
	mockEventHandler := NewEventHandler(mockEventCtr)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/events", strings.NewReader(eventJsonEdit))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockEventHandler.EditEvent(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteEvent(t *testing.T) {
	mockDB[1] = model.Event{
		Id:          1,
		Title:       "Buy food for the weekend",
		Description: "Remember the vegetables",
	}

	var eventJsonEdit = `{
		"Id": 1,
		"Title": "Buy dog food for the weekend",
		"Description": "Remember the vegetables"
		}`
	mockEventRepo := database.NewEventRepository(mockDB)
	mockEventCtr := controller.NewEventController(mockEventRepo)
	mockEventHandler := NewEventHandler(mockEventCtr)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/events", strings.NewReader(eventJsonEdit))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockEventHandler.DeleteEvent(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteEventEventNotExist(t *testing.T) {
	var eventJsonEdit = `{
		"Id": 1,
		"Title": "Buy dog food for the weekend",
		"Description": "Remember the vegetables"
		}`
	mockEventRepo := database.NewEventRepository(mockDB)
	mockEventCtr := controller.NewEventController(mockEventRepo)
	mockEventHandler := NewEventHandler(mockEventCtr)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/events", strings.NewReader(eventJsonEdit))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockEventHandler.DeleteEvent(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
