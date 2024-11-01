package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/brentgroves/user-info-proto/app/proto"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedUsrServer
}

func (*server) GetUser(ctx context.Context, in *proto.UserRequest) (*proto.UserResponse, error) {

	others := make(map[string]string)
	others["secondary"] = "233453"
	phone := &proto.PhoneNumber{Primary: "1234567890", Others: others}
	user := &proto.User{Name: "Digvijay", Age: 23, Address: &proto.Address{Street: "Pune", City: "Pune", State: "MAHARASHTRA", Zip: "201223"}, Phone: phone, UpdatedAt: time.Now().UTC().String(), CreatedAt: time.Now().UTC().String()}
	return &proto.UserResponse{User: user, Status: 200, Error: ""}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	grpcServer := &server{}
	proto.RegisterUsrServer(s, grpcServer)
	log.Printf("Starting server on port :%v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
