package repository

import "github.com/joaquincamara/golangPractices/pkg/model"

type IEventRepository interface {
	Insert(event model.Event) error
	Update(event model.Event) error
	Delete(event model.Event) error
	FindAll() (map[int32]model.Event, error)
}
