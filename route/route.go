package route

import (
	"io"
	"log"
	"strconv"

	"golang.org/x/net/context"
)

type Server struct {
	x int
	y int
}

var mallTerminal = Place{"Mall Terminal", 7, 10}
var mainStreetS = Place{"Main Street S", 21, 5}
var secondStreetE = Place{"Second Street E", 47, 40}
var cityHall = Place{"City Hall", 25, 25}
var unionTerminal = Place{"Union Terminal", 17, 48}
var secondStreetW = Place{"Second Street W", 3, 40}
var thirdStreetW = Place{"Third Street W", 6, 21}
var busAStops = []Place{mallTerminal, mainStreetS, cityHall, secondStreetE, unionTerminal}
var busBStops = []Place{unionTerminal, secondStreetW, thirdStreetW, mallTerminal}
var busA = &BusRoute{"Temp", "Mall Terminal", "Union Terminal", busAStops}
var busB = &BusRoute{"Temp2", "Union Terminal", "Mall Terminal", busBStops}


//Set the start location given from the client
func (s *Server) SetStart(stream RouteService_SetStartServer) error {
	log.Printf("Recevied request from client to save start location. Saving...")
	var counter int
	for {
		cord, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Successfully saved starting location from client: (%d, %d)", s.x, s.y)
			return stream.SendAndClose(&Message{Body: "Server has saved starting location"})
		}
		if err != nil {
			return err
		}
		counter++
		if counter == 1 {
			if s.x, err = strconv.Atoi(cord.Body); err != nil {
				log.Fatalf("Failed to convert string to int: %s", err)
			}
		} else {
			if s.y, err = strconv.Atoi(cord.Body); err != nil {
				log.Fatalf("Failed to convert string to int: %s", err)
			}
		}
	}
}

//Find a bus from the start location given the end location
func (s *Server) FindBus(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client, checking for stop: %s", in.Body)
	if busA.checkStop(in.Body) {
		return &Message{Body: busA.name}, nil
	}
	return &Message{Body: "Couldn't find a bus for your needs"}, nil
}
