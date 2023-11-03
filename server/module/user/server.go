package users

import pb "gRPC-go/proto"

type UserServer struct{
	pb.UserServiceServer
}