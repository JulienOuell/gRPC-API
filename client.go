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

func getClientMessage() (string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Give the server a message: ")
	message, _ := reader.ReadString('\n')

	return message
}

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":%d", 9000), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := route.NewRouteServiceClient(conn)

	message := getClientMessage()

	response, err := c.SayHello(context.Background(), &route.Message{Body: message})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
