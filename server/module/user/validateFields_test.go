package users

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestValidateId(t *testing.T) {
	InitDB()
	err := ValidateId(1)
	assert.NoError(t, err)
}

func TestValidateIdWhenInvalid(t *testing.T) {
	InitDB()

	err := ValidateId(0)
	assert.Error(t, err)
}

func TestValidateFname(t *testing.T) {
	err := ValidateFname("John")
	assert.NoError(t, err)
}

func TestValidateFnameWhenLengthLessThanMinLength(t *testing.T) {
	err := ValidateFname("A")
	assert.Error(t, err)
}

func TestValidateFnameWhenLengthGreaterThanMaxLength(t *testing.T) {
	longName := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
	err := ValidateFname(longName)
	assert.Error(t, err)
}

func TestValidateFnameWhenInvalidCharacters(t *testing.T) {
	err := ValidateFname("John123")
	assert.Error(t, err)
}

func TestValidateCity(t *testing.T) {
	err := ValidateCity("New York")
	assert.NoError(t, err)
}

func TestValidateCityWhenInvalidCharacter(t *testing.T) {
	err := ValidateCity("Los Angeles!")
	assert.Error(t, err)
}

func TestValidateCityWhenLengthLessThanMinLength(t *testing.T) {
	err := ValidateCity("A")
	assert.Error(t, err)
}

func TestValidateCityWhenLengthGreaterThanMaxLength(t *testing.T) {
	longCity := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
	err := ValidateCity(longCity)
	assert.Error(t, err)
}

func TestValidatePhone(t *testing.T) {
	err := ValidatePhone(1234567890)
	assert.NoError(t, err)
}

func TestValidatePhoneWhenLengthLessThanMinLength(t *testing.T) {
	err := ValidatePhone(12345)
	assert.Error(t, err)
}

func TestValidatePhoneWhenLengthGreaterThanMaxLength(t *testing.T) {
	err := ValidatePhone(12345678901)
	assert.Error(t, err)
}

func TestValidateHeight(t *testing.T) {
	err := ValidateHeight(5.8)
	assert.NoError(t, err)

	err = ValidateHeight(0.5)
	assert.Error(t, err)
}
