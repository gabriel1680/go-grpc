package services

import (
	"context"
	"fmt"

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
