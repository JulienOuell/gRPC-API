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

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := route.Server{}

	grpcServer := grpc.NewServer()

	route.RegisterRouteServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
