package users

import (
	pb "gRPC-go/proto"
)

func GenerateRequest(id int32) *pb.UserIdRequest {
    return &pb.UserIdRequest{Id: id}
}