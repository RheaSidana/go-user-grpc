package users

import (
	"context"
	"errors"
	pb "gRPC-go/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func(s *UserServer) GetUserDetails(c context.Context, req *pb.UserIdRequest) (*pb.User, error){
	if err := validateGetUserDetailsRequest(req); err != nil {
		return nil, err
	} 
	
	id := req.Id

	user, found := DB.Users[id]

	violations := validateUser(user)
	if violations != nil {
		return nil, allViolations(violations)
	}

    if !found {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return user, nil
}

func validateGetUserDetailsRequest(req *pb.UserIdRequest) error {
	if err := ValidateId(req.GetId()); err != nil {
		return errors.New(fieldViolation("id", err))
	}

	return nil
}