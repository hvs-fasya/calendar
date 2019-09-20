package models

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	pb "github.com/hvs-fasya/calendar/calendar"
)

var CalendarProtoAdapter = &ProtoAdapter{}

type ProtoAdapter struct {
	ToProto
	FromProto
}
type ToProto interface {
	EventToProto(Event) (pb.Event, bool)
	EventsToProto(Events) (pb.Events, bool)
}
type FromProto interface {
	ProtoToEvent(pb.Event) (Event, bool)
	ProtoToEvents(pb.Events) (Events, bool)
}

func (a *ProtoAdapter) EventToProto(e Event) (event pb.Event, ok bool) {
	ts, err := ptypes.TimestampProto(e.TimeStamp)
	if err != nil {
		e := errors.Wrap(err, "EventToProto")
		zap.L().Error(e.Error())
		return
	}
	event = pb.Event{
		UUID: e.UUID,
		Data: &pb.EventData{
			Title:        e.Title,
			TimeStamp:    ts,
			Duration:     ptypes.DurationProto(e.Duration),
			Description:  e.Description,
			NotifyBefore: ptypes.DurationProto(*e.NotifyBefore),
		},
		User: &pb.User{},
	}
	ok = true
	return
}
func (a *ProtoAdapter) ProtoToEvent(e pb.Event) (event Event, ok bool) {
	var user = User{}
	d, err := ptypes.Duration(e.Data.Duration)
	if err != nil {
		e := errors.Wrap(err, "ProtoToEvent")
		zap.L().Error(e.Error())
		return
	}
	ts, err := ptypes.Timestamp(e.Data.TimeStamp)
	if err != nil {
		e := errors.Wrap(err, "ProtoToEvent")
		zap.L().Error(e.Error())
		return
	}
	notifyDuration, err := ptypes.Duration(e.Data.NotifyBefore)
	if err != nil {
		e := errors.Wrap(err, "ProtoToEvent")
		zap.L().Error(e.Error())
		return
	}
	event = Event{
		UUID: e.UUID,
		EventData: EventData{
			Title:        e.Data.Title,
			TimeStamp:    ts,
			Duration:     d,
			Description:  e.Data.Description,
			NotifyBefore: &notifyDuration,
		},
		User: user,
	}
	ok = true
	return
}

func (a *ProtoAdapter) EventsToProto(ee Events) (events pb.Events, ok bool) {
	for k, e := range ee {
		event, ok := a.EventToProto(e)
		if !ok {
			return
		}
		events.Events[k] = &event
	}
	ok = true
	return
}
func (a *ProtoAdapter) ProtoToEvents(pb.Events) (Events, bool) {
	return Events{}, true
}
