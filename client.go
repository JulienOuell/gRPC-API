package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/JulienOuell/gRPC-API/route"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":%d", 9000), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := route.NewRouteServiceClient(conn)

	response, err := c.SayHello(context.Background(), &route.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
