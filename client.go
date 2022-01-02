package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/JulienOuell/gRPC-API/route"
)

//Function to get user input
func getClientMessage(text string) (string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(text)
	message, _ := reader.ReadString('\n')
	message = strings.TrimSuffix(message, "\n")

	return message
}

func setStart(c route.RouteServiceClient) {
	var points []*route.Message
	var point string
	for i := 0; i < 2; i++ {
		point = getClientMessage("Provide starting x/y location (x first): ")
		message := &route.Message{Body: point}
		points = append(points, message)
	}

	stream, err := c.SetStart(context.Background())
	if err != nil {
		log.Fatalf("Error when starting stream: %v", err)
	}
	for _, point := range points {
		if err := stream.Send(point); err != nil {
			log.Fatalf("Error when sending point to server: %v", err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error when closing stream: %v", err)
	}
	log.Printf(reply.Body)
}

func main() {

	//Client connection to port 9000 so we can talk with the server
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":%d", 9000), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %s", err)
	}
	defer conn.Close()

	c := route.NewRouteServiceClient(conn)

	setStart(c)
	message := getClientMessage("Give stop: ")

	//Send message to server stub and get response
	response, err := c.FindBus(context.Background(), &route.Message{Body: message})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
