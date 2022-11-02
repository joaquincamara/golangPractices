package event

type IEventController interface {
	Add(event Event) error
	Edit(event Event) error
	Delete(event Event) error
	FindAll() ([]Event, error)
}

type eventController struct {
	eventRepo IEventRepository
}

func NewEventController(eventRepo IEventRepository) eventController {
	return eventController{eventRepo: eventRepo}
}

func (ec *eventController) Add(e Event) error {
	if err := ec.eventRepo.Insert(e); err != nil {
		return err
	}

	return nil
}

func (ec *eventController) Edit(e Event) error {
	if err := ec.eventRepo.Update(e); err != nil {
		return err
	}

	return nil
}

func (ec *eventController) Delete(e Event) error {
	if err := ec.eventRepo.Update(e); err != nil {
		return err
	}

	return nil
}

func (ec *eventController) FindAll() ([]Event, error) {
	elements, _ := ec.eventRepo.FindAll()
	list := make([]Event, 0, len(elements))
	for _, v := range elements {
		list = append(list, v)
	}

	return list, nil
}
