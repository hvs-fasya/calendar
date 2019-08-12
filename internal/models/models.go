package models

import "time"

//Event event structure
type Event struct {
	UUID string
	EventData
	User
}

//EventData event attributes
type EventData struct {
	Title        string
	TimeStamp    time.Time
	Duration     time.Duration
	Description  string
	NotifyBefore *time.Duration
}

//User user structure
type User struct{}
