package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/JulienOuell/gRPC-API/route"
)

func main() {

	fmt.Println("Server has been started!")

	//Create connection for port 9000
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//Create service
	s := route.Server{}

	//New grpc server
	grpcServer := grpc.NewServer()

	//Add service to the grpc server
	route.RegisterRouteServiceServer(grpcServer, &s)

	//grpc server now listens on port 9000
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
