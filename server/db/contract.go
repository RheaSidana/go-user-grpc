package db

import (
	pb "gRPC-go/proto"
)

type Db struct {
    Users map[int32]*pb.User
}