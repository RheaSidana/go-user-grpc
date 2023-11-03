package users

import (
	"context"
	"fmt"
	pb "gRPC-go/proto"
	"log"
)

func CallGetUserDetails(userRequest *pb.UserIdRequest) {
	user, err := Client.GetUserDetails(context.Background(), userRequest)
	if err != nil {
		log.Printf("\n Failed to retrieve user details: %v", err)
		return
	}

	userStr := fmt.Sprintf("\nUser ID: %d\nName: %s\nCity: %s\nPhone: %d\nHeight: %.2f\nMarried: %v\n\n",
		user.Id, user.Fname, user.City, user.Phone, user.Height, user.Married)

	log.Printf("User Details: %+v", userStr)
}
