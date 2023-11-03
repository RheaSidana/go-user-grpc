package users

import (
	"testing"
	"github.com/stretchr/testify/assert"
	pb "gRPC-go/proto"
)

func TestValidateUser(t *testing.T) {
	InitDB()
	user := &pb.User{
		Id:      1,
		Fname:   "John",
		City:    "New York",
		Phone:   1234567890,
		Height:  5.8,
		Married: true,
	}
	violations := validateUser(user)
	assert.Empty(t, violations)

	invalidUser := &pb.User{
		Id:      0,
		Fname:   "123John",
		City:    "Los Angeles!",
		Phone:   12345,
		Height:  0.05,
		Married: false,
	}
	violations = validateUser(invalidUser)
	assert.NotEmpty(t, violations)
	assert.Len(t, violations, 5) 
}

func TestValidateUserWithValidValues(t *testing.T) {
	InitDB()
	user := &pb.User{
		Id:      1,
		Fname:   "John",
		City:    "New York",
		Phone:   1234567890,
		Height:  5.8,
		Married: true,
	}
	violations := validateUser(user)
	assert.Empty(t, violations)
}
