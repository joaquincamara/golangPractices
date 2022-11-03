package database

import (
	"errors"

	"github.com/joaquincamara/golangPractices/pkg/model"
	"github.com/joaquincamara/golangPractices/pkg/repository"
)

type DB map[int32]model.Event

type EventRespository struct {
	Query repository.IEventRepository
	DB    DB
}

func NewEventRepository(db DB) repository.IEventRepository {
	return &EventRespository{DB: db}
}

//EventRespository methods
func (er *EventRespository) Insert(e model.Event) error {

	if _, ok := er.DB[e.Id]; ok {
		return errors.New("event already exists")
	}

	er.DB[e.Id] = e
	return nil
}

func (er *EventRespository) Update(e model.Event) error {

	if _, ok := er.DB[e.Id]; !ok {
		return errors.New("event not exist")
	}

	er.DB[e.Id] = e
	return nil
}

func (er *EventRespository) Delete(e model.Event) error {

	if _, ok := er.DB[e.Id]; !ok {
		return errors.New("event not exist")
	}
	delete(er.DB, e.Id)
	return nil
}
func (er *EventRespository) FindAll() (map[int32]model.Event, error) {
	return er.DB, nil
}
