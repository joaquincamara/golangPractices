package controller

import (
	"errors"
	"testing"

	"github.com/joaquincamara/golangPractices/pkg/database"
	"github.com/joaquincamara/golangPractices/pkg/model"
	"github.com/stretchr/testify/assert"
)

var mockDB = make(database.DB)

func TestAdd(t *testing.T) {
	mockEvent := model.Event{
		Id:          1,
		Title:       "Buy dog food",
		Description: "Need to have a good flavor",
	}
	mockEventRepo := database.NewEventRepository(mockDB)
	mockEventCtr := NewEventController(mockEventRepo)

	assert.Nil(t, mockEventCtr.Add(mockEvent))
}

func TestAddExistingElement(t *testing.T) {
	mockDB[1] = model.Event{
		Id:          1,
		Title:       "Buy dog food",
		Description: "Need to have a good flavor",
	}
	mockEventRepo := database.NewEventRepository(mockDB)
	mockEventCtr := NewEventController(mockEventRepo)

	assert.Equal(t, mockEventCtr.Add(mockDB[1]), errors.New("event already exists"))
}

func TestEdit(t *testing.T) {
	mockDB[1] = model.Event{
		Id:          1,
		Title:       "Buy dog food",
		Description: "Need to have a good flavor",
	}
	mockEventEdtion := model.Event{
		Id:          1,
		Title:       "Buy human food",
		Description: "Need to have a good flavor",
	}
	mockEventRepo := database.NewEventRepository(mockDB)
	mockEventCtr := NewEventController(mockEventRepo)
	assert.Nil(t, mockEventCtr.Edit(mockEventEdtion))
}

func TestEditEventNotExist(t *testing.T) {
	mockEventEdtion := model.Event{
		Id:          1,
		Title:       "Buy human food",
		Description: "Need to have a good flavor",
	}
	mockEventRepo := database.NewEventRepository(mockDB)
	mockEventCtr := NewEventController(mockEventRepo)
	assert.Equal(t, mockEventCtr.Edit(mockEventEdtion), errors.New("event not exist"))
}

func TestDelete(t *testing.T) {
	mockDB[1] = model.Event{
		Id:          1,
		Title:       "Buy human food",
		Description: "Need to have a good flavor",
	}
	mockEventEdtion := model.Event{
		Id:          1,
		Title:       "Buy human food",
		Description: "Need to have a good flavor",
	}
	mockEventRepo := database.NewEventRepository(mockDB)
	mockEventCtr := NewEventController(mockEventRepo)
	assert.Nil(t, mockEventCtr.Delete(mockEventEdtion))
}

func TestDeleteEventNotExist(t *testing.T) {
	mockEventEdtion := model.Event{
		Id:          1,
		Title:       "Buy human food",
		Description: "Need to have a good flavor",
	}
	mockEventRepo := database.NewEventRepository(mockDB)
	mockEventCtr := NewEventController(mockEventRepo)
	assert.Equal(t, mockEventCtr.Delete(mockEventEdtion), errors.New("event not exist"))
}
