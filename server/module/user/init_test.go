package users

import (
	"fmt"
	"gRPC-go/server/initializer"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestRegisterUserServer(t *testing.T) {
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	defer server.Stop()

	initializer.Server = server
	initializer.Lis = listener

	RegisterUserServer()
	defer initializer.Lis.Close()

	services := server.GetServiceInfo()
	fmt.Println(services)

	_, ok := services["users.UserService"]
	require.True(t, ok, "users.UserService service should be registered")
}

