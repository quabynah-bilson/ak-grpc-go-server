package main

import (
	"fmt"
	pb "github.com/quabynah-bilson/grpc-go/protos"
	"github.com/quabynah-bilson/grpc-go/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = 8080
)

func main() {
	// 1. create an instance of your server
	app := grpc.NewServer()
	
	// 2. add routes (services)
	pb.RegisterEchoServiceServer(app, services.NewEchoService())
	
	// enable reflection
	reflection.Register(app)
	
	// 3.set up a listener for the server on the specified port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", port, err)
	}
	
	// listen for connections on that network
	if err = app.Serve(lis); err != nil {
		log.Fatalf("failed to serve on port %d: %v", port, err)
	}
}
