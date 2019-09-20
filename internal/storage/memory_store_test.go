package storage

import (
	"reflect"
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/hvs-fasya/calendar/internal/models"
	"github.com/hvs-fasya/calendar/internal/test_utils"
)

func TestMemStore_GetEvents(t *testing.T) {
	type args struct {
		event *models.Event
	}
	var existEvent = test_utils.NewRandomEvent()
	var notExistEvent = test_utils.NewRandomEvent()
	tests := []struct {
		name    string
		wantLen int
	}{
		{"get_events", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := InitMemoryStore()
			s.CreateEvent(existEvent)
			got := s.GetEvents()
			if len(s.events) != tt.wantLen {
				t.Errorf("GetEvents() events store len = %v, want %v", len(s.events), tt.wantLen)
				return
			}
			var found bool
			var foundNotExist bool
			for _, e := range got {
				if reflect.DeepEqual(e.EventData, existEvent.EventData) {
					found = true
				}
				if reflect.DeepEqual(e.EventData, notExistEvent.EventData) {
					foundNotExist = true
				}
			}
			if !found {
				t.Errorf("GetEvents() do not contain created event")
			}
			if foundNotExist {
				t.Errorf("GetEvents() contains not created event")
			}
		})
	}
}

func TestMemStore_CreateEvent(t *testing.T) {
	type args struct {
		event models.Event
	}
	var randEvent = test_utils.NewRandomEvent()

	t.Run("create_event", func(t *testing.T) {
		s := InitMemoryStore()
		got, err := s.CreateEvent(randEvent)
		if err != nil {
			t.Errorf("CreateEvent() error = %v, want no error", err)
			return
		}
		if !reflect.DeepEqual(got.EventData, randEvent.EventData) {
			t.Errorf("CreateEvent() got = %v, want %v", got, randEvent.EventData)
		}
		created, err := s.GetEventByID(got.UUID)
		if err != nil {
			t.Errorf("CreateEvent() error = %v, want no error", err)
			return
		}
		if !reflect.DeepEqual(created.EventData, randEvent.EventData) {
			t.Errorf("CreateEvent() storage does not contain event: %v", randEvent)
		}
	})
}

func TestMemStore_UpdateEvent(t *testing.T) {
	var existEvent = test_utils.NewRandomEvent()
	var updateData = test_utils.NewRandomEvent().EventData
	s := InitMemoryStore()
	exist, _ := s.CreateEvent(existEvent)

	t.Run("update_event", func(t *testing.T) {
		updated, err := s.UpdateEvent(models.Event{UUID: exist.UUID, EventData: updateData})
		if err != nil {
			t.Errorf("UpdateEvent() error = %v, want no error", err)
			return
		}
		if !reflect.DeepEqual(updated.EventData, updateData) {
			t.Errorf("UpdateEvent() updated event data = %v, want %v", updated.EventData, updateData)
		}
		stored, err := s.GetEventByID(exist.UUID)
		if err != nil {
			t.Errorf("UpdateEvent() error = %v, want no error", err)
		}
		if !reflect.DeepEqual(stored.EventData, updateData) {
			t.Errorf("UpdateEvent() stored event data = %v, want %v", stored.EventData, updateData)
		}
	})
}

func TestMemStore_DeleteEvent(t *testing.T) {
	var existEvent = test_utils.NewRandomEvent()
	var notFoundUUID = uuid.NewV4().String()
	s := InitMemoryStore()
	exist, _ := s.CreateEvent(existEvent)

	t.Run("delete_event", func(t *testing.T) {
		if err := s.DeleteEvent(exist.UUID); err != nil {
			t.Errorf("DeleteEvent() error = %v, want no error", err)
			return
		}
		_, err := s.GetEventByID(exist.UUID)
		if err == nil {
			t.Errorf("DeleteEvent() store contains deleted event")
		}
	})
	t.Run("delete_event_not_found", func(t *testing.T) {
		if err := s.DeleteEvent(notFoundUUID); err == nil {
			t.Errorf("DeleteEvent() should return error for not exist uuid")
			return
		}
	})
}
