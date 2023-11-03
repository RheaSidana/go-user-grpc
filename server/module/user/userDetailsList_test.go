package users

import (
	"context"
	"os"
	"testing"
	"time"
	"fmt"

	pb "gRPC-go/proto"
	deferring "gRPC-go/server/defer"
	"gRPC-go/server/initializer"
	"io"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGetUserDetailsByIDs(t *testing.T) {
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
    time.Sleep(time.Second * 3)

    port := ":" + os.Getenv("PORT")
    host := os.Getenv("HOST")
    conn, err := grpc.Dial(
        host + port,
        grpc.WithTransportCredentials(
            insecure.NewCredentials(),
        ),
    )
    if err != nil {
        t.Errorf("Did not connect: %v", err)
    }
    defer conn.Close()

    userClient := pb.NewUserServiceClient(conn)

    userIds := []int32{2, 5}
    validRequest := &pb.UserIdListRequest{Id: userIds}
    stream, err := userClient.GetUserDetailsByIDs(context.Background())
    if err != nil {
        t.Errorf("Error creating stream: %v", err)
    }

    go func() {
        if err := stream.Send(validRequest); err != nil {
            t.Errorf("Error sending request: %v", err)
        }
        stream.CloseSend()
    }()

    time.Sleep(time.Second)

    for {
        usersList, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            t.Errorf("Error receiving response: %v", err)
        }

        for _, user := range usersList.User {
			userStr := getUser(user)
			orgStr := getUser(DB.Users[user.Id])
            assert.Equal(t, orgStr, userStr)
        }
    }

    os.Unsetenv("HOST")
    os.Unsetenv("NETWORK")
    os.Unsetenv("PORT")
}

func getUser(user *pb.User) string {
	return fmt.Sprintf("\nUser ID: %d\nName: %s\nCity: %s\nPhone: %d\nHeight: %.2f\nMarried: %v\n\n",
	user.Id, user.Fname, user.City, user.Phone, user.Height, user.Married)
}