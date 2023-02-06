package main

import (
	"context"
	"fmt"
	"log"
	"time"
	pb "usermgmt/usermgmt"

	grpc "google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	fmt.Println("Starting User Management Service")

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed in connection %v", err)
	}
	defer conn.Close()

	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	var new_user = make(map[string]int32)

	new_user["Roushan"] = 29
	new_user["Kanchan"] = 30

	for name, age := range new_user {

		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("Could not create user %v", err)
		}

		log.Printf(` User Details :
		User Id : %d 
		Name : %s 
		Age : %d`, r.GetId(), r.GetName(), r.GetAge())
	}

}
