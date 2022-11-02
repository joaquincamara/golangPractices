package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock json request
var eventJson = `{
	"Id": 1,
	"Title": "Buy food for the weekend",
	"Description": "Remember the vegetables"
	}`

func TestCreateEvent(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/v1/events", strings.NewReader(eventJson))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Assertions
	CreateEvent(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
