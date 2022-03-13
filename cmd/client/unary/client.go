package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gabriel1680/go-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	CreateRequest(client)
}

func CreateRequest(client pb.UserServiceClient) {
	req := &pb.UserRequest{
		FirstName: "Gabriel",
		LastName:  "Lopes",
		Email:     "gabriel.lopes123456789@gmail.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not send the request: %v", err)
	}

	fmt.Println(res)
}
