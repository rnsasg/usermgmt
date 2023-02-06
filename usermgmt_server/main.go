package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	pb "usermgmt/usermgmt"

	grpc "google.golang.org/grpc"
)

const (
	port = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Receviced a reuest from %s", in.GetName())
	var user_id int32 = int32(rand.Intn(1000))
	return &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}, nil
}

func main() {

	fmt.Println("Starting User Management Service Server")

	// Start Listening on port
	lis, _ := net.Listen("tcp", port)

	// Start a gRPC server

	s := grpc.NewServer()

	pb.RegisterUserManagementServer(s, &UserManagementServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed in serving %v", err)
	}
}
