package users

import (
	"fmt"
	"gRPC-go/server/module/validator"
	"regexp"
)

var (
	isValidCity = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`).MatchString
	isValidFname = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)


func ValidateId(value int32) error {
	return validator.ValidateInt32(value, 1, int32(len(DB.Users)))
}

func ValidateFname(value string) error {
	if err := validator.ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidFname(value) {
		return fmt.Errorf("must contain only letters or spaces")
	}
	return nil
}

func ValidateCity(value string) error {
	if err := validator.ValidateString(value, 2, 100); err != nil {
		return err
	}
	if !isValidCity(value) {
		return fmt.Errorf("must contain only letters, numbers or spaces")
	}
	return nil
}

func ValidatePhone(value int64) error {
	return validator.ValidateInt64(value, 0000000000, 10000000000)
}

func ValidateHeight(value float64) error {
	return validator.ValidateDouble(value, 1.0, 2)
}
