package users

import (
	"errors"
	pb "gRPC-go/proto"
	"io"
	"log"
	"sync"
)

func (s *UserServer) GetUserDetailsByIDs(stream pb.UserService_GetUserDetailsByIDsServer) error {
	for {
		userID, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		wg := sync.WaitGroup{}

		for _, userID := range userID.Id {
			if err := validateGetUserDetailsByIDsRequest(userID); err != nil {
				return err
			}
			wg.Add(1)

			go func(id int32) {
				defer wg.Done()

				user, found := DB.Users[id]

				if !found {
					log.Printf("User not found")
					return
				}

				if err := stream.Send(
					&pb.UsersList{User: []*pb.User{user}},
				); err != nil {					
					return
				}
			}(userID)
		}
		wg.Wait()
	}

	return nil
}

func validateGetUserDetailsByIDsRequest(id int32) error {
	if err := ValidateId(id); err != nil {
		return errors.New(fieldViolation("id", err))
	}

	return nil
}