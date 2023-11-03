package users

import (
	pb "gRPC-go/proto"
)

func validateUser(user *pb.User)(violations []string) {
	if err := ValidateId(user.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}

	if err := ValidateFname(user.GetFname()); err != nil {
		violations = append(violations, fieldViolation("fname", err))
	}

	if err := ValidateCity(user.GetCity()); err != nil {
		violations = append(violations, fieldViolation("city", err))
	}
	
	if err := ValidatePhone(user.GetPhone()); err != nil {
		violations = append(violations, fieldViolation("phone", err))
	}
	
	if err := ValidateHeight(user.GetHeight()); err != nil {
		violations = append(violations, fieldViolation("height", err))
	}

	return violations
}