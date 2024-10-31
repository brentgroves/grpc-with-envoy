package main

import (
 "context"
 "log"
 "net"
 "google.golang.org/grpc"
)

// YourService should implement the methods of your GRPC service.
type YourService struct{}

// Implement your GRPC methods here.

func main() {
 listener, err := net.Listen("tcp", ":8080")
 if err != nil {
  log.Fatalf("Failed to listen: %v", err)
 }

 grpcServer := grpc.NewServer()
 // Register your service here.
 
 if err := grpcServer.Serve(listener); err != nil {
  log.Fatalf("Failed to serve: %v", err)
 }
