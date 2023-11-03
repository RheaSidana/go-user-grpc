package initializer

import (
	"log"
	"os"
	"net"

	"google.golang.org/grpc"
)

var Server *grpc.Server
var Lis net.Listener

func CreateNewGrpcServer() {
	protocol := os.Getenv("NETWORK")
	port := ":" + os.Getenv("PORT")
	lis, err := net.Listen(protocol, port)
	if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

	Server = grpc.NewServer()
	Lis = lis
}