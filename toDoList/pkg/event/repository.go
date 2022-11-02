package event

type IEventRepository interface {
	Insert(event Event) error
	Update(event Event) error
	Delete(event Event) error
	FindAll() (map[int32]Event, error)
}
