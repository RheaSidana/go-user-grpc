package users

import (
	"gRPC-go/client/initializer"
	pb "gRPC-go/proto"
)

var Client pb.UserServiceClient

func NewUserClient() {
	Client = pb.NewUserServiceClient(initializer.Conn)
}