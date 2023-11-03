package users

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "gRPC-go/proto"
)

func CallGetUserDetailsByIDs(userIds []int32) {
	stream, err := Client.GetUserDetailsByIDs(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	go func() {
		request := &pb.UserIdListRequest{Id: userIds}
		if err := stream.Send(request); err != nil {
			log.Fatalf("Error sending request: %v", err)
		}
		stream.CloseSend()
	}()

	for {
		usersList, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Error receiving response: %v", err)
		}

		for _, user := range usersList.User {
			userStr := fmt.Sprintf("\nUser ID: %d\nName: %s\nCity: %s\nPhone: %d\nHeight: %.2f\nMarried: %v\n\n",
				user.Id, user.Fname, user.City, user.Phone, user.Height, user.Married)
			log.Printf("User Details: %+v\n", userStr)
		}
	}
}
