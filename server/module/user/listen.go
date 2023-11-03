package users

import (
	"gRPC-go/server/initializer"
	"log"
)

func ListenUserServer() {
	if err := initializer.Server.Serve(initializer.Lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}