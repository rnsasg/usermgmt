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
	user_list *pb.UserList
}

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{
		user_list: &pb.UserList{},
	}
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Receviced a reuest from %s", in.GetName())
	var user_id int32 = int32(rand.Intn(1000))
	created_user := &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}

	s.user_list.Users = append(s.user_list.Users, created_user)

	return created_user, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context, in *pb.GetUserParams) (*pb.UserList, error) {
	return s.user_list, nil
}

func (server *UserManagementServer) Run() error {

	// Start Listening on port
	lis, _ := net.Listen("tcp", port)

	// Start a gRPC server
	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, server)

	return s.Serve(lis)
}

func main() {

	fmt.Println("Starting User Management Service Server")
	var usr_mgmt_srv *UserManagementServer = NewUserManagementServer()
	if err := usr_mgmt_srv.Run(); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
