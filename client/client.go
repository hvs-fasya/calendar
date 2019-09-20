package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	pb "github.com/hvs-fasya/calendar/calendar"
)

var (
	dialOpts   []grpc.DialOption
	serverAddr string

	Client pb.CalendarClient
)

type Cfg struct {
	Host string
	Port string
}

func Run(cfg *Cfg) error {
	serverAddr = cfg.Host + ":" + cfg.Port
	dialOpts = append(dialOpts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, dialOpts...)
	if err != nil {
		return err
	}
	defer conn.Close()
	Client = pb.NewCalendarClient(conn)
	getEvents()
	return nil
}

func getEvents() (events *pb.Events, err error) {
	events, err = Client.GetEvents(context.Background(), &pb.GetEventsParams{})
	if err != nil {
		return
	}
	fmt.Printf("EVENTS: %+v\n", events)
	return
}
