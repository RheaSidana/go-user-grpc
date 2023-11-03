package users

import (
	"context"
	"os"
	"testing"
	"time"

	pb "gRPC-go/proto"
	deferring "gRPC-go/server/defer"
	"gRPC-go/server/initializer"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetUserDetails(t *testing.T) {
	os.Setenv("NETWORK", "tcp")
	os.Setenv("PORT", "50051")
	os.Setenv("HOST", "localhost")

	initializer.CreateNewGrpcServer()
	RegisterUserServer()
	defer deferring.DeferServer()
	go func() {
		ListenUserServer()
		defer deferring.DeferListener()
	}()
	time.Sleep(time.Second*3)

	port := ":" + os.Getenv("PORT")
	host := os.Getenv("HOST")
	conn, err := grpc.Dial(
		host+port,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		t.Errorf("Did not connect: %v", err)
	}
	defer conn.Close()

	userClient := pb.NewUserServiceClient(conn)

	validRequest := &pb.UserIdRequest{Id: 1}
	user, err := userClient.GetUserDetails(context.Background(), validRequest)
	require.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Steve", user.Fname)
	assert.Equal(t, "LA", user.City)
	assert.Equal(t, int64(1234567890), user.Phone)
	assert.Equal(t, float64(5.8), user.Height)
	assert.True(t, user.Married)
	
	os.Unsetenv("HOST")
	os.Unsetenv("NETWORK")
	os.Unsetenv("PORT")
}

func TestGetUserDetailsWhenIdDoesNotExist(t *testing.T) {
	os.Setenv("NETWORK", "tcp")
	os.Setenv("PORT", "50051")
	os.Setenv("HOST", "localhost")

	initializer.CreateNewGrpcServer()
	RegisterUserServer()
	defer deferring.DeferServer()
	go func() {
		ListenUserServer()
		defer deferring.DeferListener()
	}()
	time.Sleep(time.Second*3)

	port := ":" + os.Getenv("PORT")
	host := os.Getenv("HOST")
	conn, err := grpc.Dial(
		host+port,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		t.Errorf("Did not connect: %v", err)
	}
	defer conn.Close()

	userClient := pb.NewUserServiceClient(conn)

	invalidRequest := &pb.UserIdRequest{Id: 1000}
	_, err = userClient.GetUserDetails(context.Background(), invalidRequest)
	require.Error(t, err)
	st, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.Unknown, st.Code())
		
	os.Unsetenv("HOST")
	os.Unsetenv("NETWORK")
	os.Unsetenv("PORT")
}

func TestGetUserDetailsWhenInvalidId(t *testing.T) {
	os.Setenv("NETWORK", "tcp")
	os.Setenv("PORT", "50051")
	os.Setenv("HOST", "localhost")

	initializer.CreateNewGrpcServer()
	RegisterUserServer()
	defer deferring.DeferServer()
	go func() {
		ListenUserServer()
		defer deferring.DeferListener()
	}()
	time.Sleep(time.Second*3)

	port := ":" + os.Getenv("PORT")
	host := os.Getenv("HOST")
	conn, err := grpc.Dial(
		host+port,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		t.Errorf("Did not connect: %v", err)
	}
	defer conn.Close()

	userClient := pb.NewUserServiceClient(conn)

	invalidRequest := &pb.UserIdRequest{Id: -1}
	_, err = userClient.GetUserDetails(context.Background(), invalidRequest)
	require.Error(t, err)
	st, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.Unknown, st.Code())
		
	os.Unsetenv("HOST")
	os.Unsetenv("NETWORK")
	os.Unsetenv("PORT")
}
