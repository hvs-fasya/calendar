package storage

import (
	"sync"

	"github.com/pkg/errors"
	"github.com/satori/go.uuid"

	"github.com/hvs-fasya/calendar/internal/models"
)

//MemStore memory storage
type MemStore struct {
	sync.Mutex
	events map[string]models.Event
	users  map[string]models.User
}

//InitMemoryStore initialize in memory store
func InitMemoryStore() *MemStore {
	mem := &MemStore{}
	mem.events = make(map[string]models.Event, 0)
	mem.users = make(map[string]models.User, 0)
	return mem
}

//GetEvents get all events from store
func (s *MemStore) GetEvents() map[string]models.Event {
	events := make(map[string]models.Event)
	for k, e := range s.events {
		events[k] = e
	}
	return events
}

//CreateEvent store event to memory
func (s *MemStore) CreateEvent(event models.Event) (models.Event, error) {
	event.UUID = uuid.NewV4().String()
	s.Lock()
	s.events[event.UUID] = event
	s.Unlock()
	return event, nil
}

//UpdateEvent update already stored event
func (s *MemStore) UpdateEvent(event models.Event) (models.Event, error) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.events[event.UUID]; !ok {
		return event, errors.Wrap(ErrNotFound, "[mem store] update event")
	}
	s.events[event.UUID] = event
	return event, nil
}

//DeleteEvent rm event from store
func (s *MemStore) DeleteEvent(uuid string) error {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.events[uuid]; !ok {
		return errors.Wrap(ErrNotFound, "[mem store] delete event")
	}
	delete(s.events, uuid)
	return nil
}

//GetEventByID get event by uuid
func (s *MemStore) GetEventByID(uuid string) (models.Event, error) {
	var event models.Event
	event, ok := s.events[uuid]
	if !ok {
		return event, errors.Wrap(ErrNotFound, "[mem store] get event by id")
	}
	return event, nil
}
