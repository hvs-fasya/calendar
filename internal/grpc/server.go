package grpc

import (
	"context"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"

	pb "github.com/hvs-fasya/calendar/calendar"
	"github.com/hvs-fasya/calendar/internal/utils"
)

type Cfg struct {
	Port string
}

type calendarServer struct{}

func Run(cfg *Cfg) error {
	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterCalendarServer(s, &calendarServer{})
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (s *calendarServer) CreateEvent(context.Context, *pb.Event) (*pb.Event, error) {
	return &pb.Event{}, nil
}
func (s *calendarServer) UpdateEvent(context.Context, *pb.Event) (*pb.Event, error) {
	return &pb.Event{}, nil
}
func (s *calendarServer) GetEvents(context.Context, *pb.GetEventsParams) (*pb.Events, error) {
	title := utils.RandStringBytes(10)
	event := &pb.Event{
		UUID: title,
		Data: &pb.EventData{
			Title:       title,
			TimeStamp:   ptypes.TimestampNow(),
			Duration:    ptypes.DurationProto(time.Duration(10 * time.Minute)),
			Description: title,
		},
		User: &pb.User{},
	}
	events := &pb.Events{
		Events: map[string]*pb.Event{
			title: event,
		},
	}
	return events, nil
}
func (s *calendarServer) DeleteEvent(context.Context, *pb.DeleteEventParams) (*pb.DeleteEventOutput, error) {
	return &pb.DeleteEventOutput{}, nil
}
