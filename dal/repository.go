package dal

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"sync"
	"time"
)

type EventEntity struct {
	Id     int
	UserId uuid.UUID
	Date   time.Time
	Event  string
}

type EventRepository interface {
	CreateEvent(event *EventEntity) error
	GetEventById(id int) (*EventEntity, error)
	SaveEvent(event *EventEntity) error
	DeleteEvent(id int) error
	GetEventsBetween(start, end time.Time) []*EventEntity
}

type InMemoryEventRepository struct {
	currentId int
	events    map[int]*EventEntity
	mutex     sync.Mutex
}

var repository = InMemoryEventRepository{
	events:    make(map[int]*EventEntity),
	currentId: 1,
	mutex:     sync.Mutex{},
}

func GetEventRepository() EventRepository {
	return &repository
}

func (r *InMemoryEventRepository) GetIndex() int {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	result := r.currentId
	r.currentId++
	return result
}

func (r *InMemoryEventRepository) CreateEvent(event *EventEntity) error {
	event.Id = r.GetIndex()
	r.events[event.Id] = event
	return nil
}

func (r *InMemoryEventRepository) GetEventById(id int) (*EventEntity, error) {
	if r.events[id] == nil {
		return nil, errors.New(fmt.Sprintf("Event with id %d not found", id))
	}
	return r.events[id], nil
}

func (r *InMemoryEventRepository) DeleteEvent(id int) error {
	if r.events[id] == nil {
		return errors.New(fmt.Sprintf("Event with id %d not found", id))
	}
	delete(r.events, id)

	return nil
}

func (r *InMemoryEventRepository) SaveEvent(event *EventEntity) error {
	return nil
}

func (r *InMemoryEventRepository) GetEventsBetween(start, end time.Time) []*EventEntity {
	result := make([]*EventEntity, 0)
	for _, event := range r.events {
		if (event.Date == start || event.Date.After(start)) && (event.Date == end || event.Date.Before(end)) {
			result = append(result, event)
		}
	}
	return result
}
