package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gabriel1680/go-grpc/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (*UserService) AddUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {

	// saving in database...
	fmt.Println(req.Email)

	return &pb.UserResponse{
		Id:       "someBigAndUniqueId",
		FullName: req.GetFirstName() + " " + req.GetLastName(),
		Email:    req.GetEmail(),
	}, nil
}

func (*UserService) AddUserWithResponseStream(req *pb.UserRequest, stream pb.UserService_AddUserWithResponseStreamServer) error {

	stream.Send(&pb.UserResponseStream{
		Status: "Started",
		User:   &pb.UserResponse{},
	})

	time.Sleep(time.Second * 2)

	stream.Send(&pb.UserResponseStream{
		Status: "Making some process with data",
		User:   &pb.UserResponse{},
	})

	time.Sleep(time.Second * 2)

	stream.Send(&pb.UserResponseStream{
		Status: "Completed",
		User: &pb.UserResponse{
			Id:       "1234",
			FullName: req.GetFirstName() + " " + req.GetLastName(),
			Email:    req.GetEmail(),
		},
	})

	return nil
}

func (*UserService) AddUsersWithRequestStream(stream pb.UserService_AddUsersWithRequestStreamServer) error {
	users := []*pb.UserResponse{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MultiUserResponse{
				Users: users,
			})
		}

		if err != nil {
			log.Fatalf("Error on receiving stream request: %v", err)
		}

		users = append(users, &pb.UserResponse{
			Id:       "1234",
			FullName: req.GetFirstName() + " " + req.GetLastName(),
			Email:    req.GetEmail(),
		})

		fmt.Println("Receiving", req.GetEmail())
	}
}
