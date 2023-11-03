package initializer

import (
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Conn *grpc.ClientConn

func CreateNewGrpcClient() {
	port := ":" + os.Getenv("PORT")
	host := os.Getenv("HOST")
	conn, err := grpc.Dial(
		host+port,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	Conn = conn
}