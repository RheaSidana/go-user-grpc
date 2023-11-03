package users

import (
	"context"
	"gRPC-go/server/initializer"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestListenUserServer(t *testing.T) {
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	defer server.Stop()

	initializer.Server = server
	initializer.Lis = listener

	RegisterUserServer()
	defer initializer.Lis.Close()

	go ListenUserServer()

	time.Sleep(time.Second)

	client, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial gRPC server: %v", err)
	}
	defer client.Close()
}
