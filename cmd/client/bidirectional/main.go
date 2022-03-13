package main

import (
	"context"
	"fmt"
	"io"
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

	stream, err := client.AddUsersWithBidirectionalStream(context.Background())
	if err != nil {
		log.Fatalf("Could not send the request: %v", err)
	}

	wait := make(chan int)

	go func() {
		for _, req := range requests {
			fmt.Println("Sending user: ", req)
			stream.Send(req)
			time.Sleep(time.Second * 1)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Could not receive the stream message from server: %v", err)
				break
			}
			fmt.Printf("Status: %v, User: %v\n", res.GetStatus(), res.GetUser().GetEmail())
		}
		close(wait)
	}()

	<-wait
}
