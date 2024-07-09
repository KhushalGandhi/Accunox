package main

import (
	"fmt"
	"log"
	"net"

	"UserService/proto"  // Update this import path to match your project
	"UserService/server" // Update this import path to match your project
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, &server.Server{}) // Register the server

	fmt.Println("Server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
