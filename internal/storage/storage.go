package storage

import (
	"errors"

	"github.com/hvs-fasya/calendar/internal/models"
)

//package level variables
var (
	//Store global storage variable
	Store Storage

	ErrNotFound = errors.New("not found")
)

//Storage storage interface
type Storage interface {
	GetEvents() map[string]models.Event
	CreateEvent(models.Event) (models.Event, error)
	UpdateEvent(models.Event) (models.Event, error)
	DeleteEvent(string) error
	GetEventByID(string) (models.Event, error)
}
