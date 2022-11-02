package database

import (
	"errors"
	"testing"

	"github.com/joaquincamara/golangPractices/pkg/event"
	"github.com/stretchr/testify/assert"
)

var mockDB = make(DB)

func TestInsert(t *testing.T) {
	mockEvent := event.Event{
		Id:          2,
		Title:       "Buy dog food",
		Description: "Need to have a good flavor",
	}
	mockEventRepo := NewEventRepository(mockDB)

	// Assertions
	assert.Nil(t, mockEventRepo.Insert(mockEvent))
}

func TestInsertExistingElement(t *testing.T) {
	mockEventRepo := NewEventRepository(mockDB)
	mockDB[1] = event.Event{
		Id:          1,
		Title:       "Pay electricity",
		Description: "Need to be before weekend",
	}

	assert.Equal(t, mockEventRepo.Insert(mockDB[1]), errors.New("event already exists"))
}

func TestUpdate(t *testing.T) {
	mockDB[1] = event.Event{
		Id:          1,
		Title:       "Buy dog food",
		Description: "Need to have a good flavor",
	}
	mockEventEdtion := event.Event{
		Id:          1,
		Title:       "Buy human food",
		Description: "Need to have a good flavor",
	}

	mockEventRepo := NewEventRepository(mockDB)
	assert.Nil(t, mockEventRepo.Update(mockEventEdtion))
}

func TestUpdateEventNotExist(t *testing.T) {
	mockEventEdtion := event.Event{
		Id:          1,
		Title:       "Buy human food",
		Description: "Need to have a good flavor",
	}

	mockEventRepo := NewEventRepository(mockDB)
	assert.Equal(t, mockEventRepo.Update(mockEventEdtion), errors.New("event not exist"))
}

func TestDelete(t *testing.T) {
	mockDB[1] = event.Event{
		Id:          1,
		Title:       "Buy human food",
		Description: "Need to have a good flavor",
	}
	mockEventEdtion := event.Event{
		Id:          1,
		Title:       "Buy human food",
		Description: "Need to have a good flavor",
	}

	mockEventRepo := NewEventRepository(mockDB)
	assert.Nil(t, mockEventRepo.Delete(mockEventEdtion))
}

func TestDeleteEventNotExist(t *testing.T) {
	mockEventEdtion := event.Event{
		Id:          1,
		Title:       "Buy human food",
		Description: "Need to have a good flavor",
	}

	mockEventRepo := NewEventRepository(mockDB)
	assert.Equal(t, mockEventRepo.Delete(mockEventEdtion), errors.New("event not exist"))
}
