package controller

import (
	"github.com/joaquincamara/golangPractices/pkg/model"
	"github.com/joaquincamara/golangPractices/pkg/repository"
)

type IEventController interface {
	Add(event model.Event) error
	Edit(event model.Event) error
	Delete(event model.Event) error
	FindAll() ([]model.Event, error)
}

type eventController struct {
	EventRepo repository.IEventRepository
}

func NewEventController(eventRepo repository.IEventRepository) eventController {
	return eventController{EventRepo: eventRepo}
}

func (ec eventController) Add(e model.Event) error {
	if err := ec.EventRepo.Insert(e); err != nil {
		return err
	}

	return nil
}

func (ec eventController) Edit(e model.Event) error {
	if err := ec.EventRepo.Update(e); err != nil {
		return err
	}

	return nil
}

func (ec eventController) Delete(e model.Event) error {
	if err := ec.EventRepo.Update(e); err != nil {
		return err
	}

	return nil
}

func (ec eventController) FindAll() ([]model.Event, error) {
	elements, _ := ec.EventRepo.FindAll()
	list := make([]model.Event, 0, len(elements))
	for _, v := range elements {
		list = append(list, v)
	}

	return list, nil
}
