package route

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) FindBus(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client, checking for stop: %s", in.Body)
	if busA.checkStop(in.Body) {
		return &Message{Body: busA.name}, nil
	}
	return &Message{Body: "Couldn't find a bus for your needs"}, nil
}
