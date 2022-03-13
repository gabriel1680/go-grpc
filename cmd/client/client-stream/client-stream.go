package main

import (
	"context"
	"fmt"
	"log"
	"time"

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
	requests := []*pb.UserRequest{
		&pb.UserRequest{
			FirstName: "Gabriel",
			LastName:  "Lopes",
			Email:     "gabriel.lopes123456789@gmail.com",
		},
		&pb.UserRequest{
			FirstName: "Guilherme",
			LastName:  "Calado",
			Email:     "guilherme123456789@gmail.com",
		},
		&pb.UserRequest{
			FirstName: "Julia",
			LastName:  "Bomfim",
			Email:     "julia123456789@gmail.com",
		},
	}

	stream, err := client.AddUsersWithRequestStream(context.Background())
	if err != nil {
		log.Fatalf("Could not send the request: %v", err)
	}

	for _, req := range requests {
		stream.Send(req)
		time.Sleep(time.Second * 1)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error on receive response: %v", err)
	}
	fmt.Println(res)
}
