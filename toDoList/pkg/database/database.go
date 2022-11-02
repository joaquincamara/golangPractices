package database

import (
	"errors"

	"github.com/joaquincamara/golangPractices/pkg/event"
)

type DB map[int32]event.Event

type EventRespository struct {
	Query event.IEventRepository
	DB    DB
}

func NewEventRepository(db DB) event.IEventRepository {
	return &EventRespository{DB: db}
}

//EventRespository methods
func (er *EventRespository) Insert(e event.Event) error {

	if _, ok := er.DB[e.Id]; ok {
		return errors.New("event already exists")
	}

	er.DB[e.Id] = e
	return nil
}

func (er *EventRespository) Update(e event.Event) error {

	if _, ok := er.DB[e.Id]; !ok {
		return errors.New("event not exist")
	}

	er.DB[e.Id] = e
	return nil
}

func (er *EventRespository) Delete(e event.Event) error {

	if _, ok := er.DB[e.Id]; !ok {
		return errors.New("event not exist")
	}
	delete(er.DB, e.Id)
	return nil
}
func (er *EventRespository) FindAll() (map[int32]event.Event, error) {
	return er.DB, nil
}
