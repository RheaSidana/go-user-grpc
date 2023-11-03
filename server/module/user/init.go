package users

import (
	"gRPC-go/server/initializer"
	pb "gRPC-go/proto"
	"log"
)

func RegisterUserServer(){
	pb.RegisterUserServiceServer(initializer.Server, &UserServer{})

	log.Printf("Server started at %v", initializer.Lis.Addr())

	InitDB()
}