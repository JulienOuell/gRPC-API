package main

import (
	"fmt"
	"log"
	"os"
	"bufio"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/JulienOuell/gRPC-API/route"
)

//Function to get user input
func getClientMessage() (string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Give the server a message: ")
	message, _ := reader.ReadString('\n')

	return message
}

func main() {

	//Client connection to port 9000 so we can talk with the server
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":%d", 9000), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := route.NewRouteServiceClient(conn)

	message := getClientMessage()

	//Send message to server stub and get response
	response, err := c.FindBus(context.Background(), &route.Message{Body: message})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
