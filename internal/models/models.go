package models

import "time"

//Event event structure
type Event struct {
	UUID      string `json:"uuid"`
	EventData `json:"data"`
	User      `json:"user"`
}

type Events map[string]Event

//EventData event attributes
type EventData struct {
	Title        string         `json:"title"`
	TimeStamp    time.Time      `json:"time_stamp"`
	Duration     time.Duration  `json:"duration"`
	Description  string         `json:"description"`
	NotifyBefore *time.Duration `json:"notify_before"`
}

//User user structure
type User struct{}
