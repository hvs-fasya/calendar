package test_utils

import (
	"time"

	"github.com/hvs-fasya/calendar/internal/models"
	"github.com/hvs-fasya/calendar/internal/utils"
)

func NewRandomEvent() models.Event {
	d := time.Duration(1 * time.Hour)
	nBefore := time.Duration(15 * time.Minute)
	return models.Event{
		EventData: models.EventData{
			Title:        utils.RandStringBytes(10),
			TimeStamp:    time.Now(),
			Duration:     d,
			Description:  utils.RandStringBytes(10),
			NotifyBefore: &nBefore,
		},
		User: models.User{},
	}
}
