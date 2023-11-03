package users

import (
	"errors"
	"testing"
)

func TestFieldViolation(t *testing.T) {
	field := "Field1"
	err := errors.New("Error message")

	result := fieldViolation(field, err)
	expectedResult := "{ Field: Field1, \n Description: Error message, \n}\n"

	if result != expectedResult {
		t.Errorf("fieldViolation(%s, %v) = %s; want %s", field, err, result, expectedResult)
	}
}

func TestAllViolations(t *testing.T) {
	violations := []string{
		"Violation 1",
		"Violation 2",
		"Violation 3",
	}

	result := allViolations(violations)
	expectedResult := "Violation 1Violation 2Violation 3"

	if result.Error() != expectedResult {
		t.Errorf("allViolations(%v) = %s; want %s", violations, result, expectedResult)
	}
}

func TestAllViolationsWhenNoViolation(t *testing.T) {
	emptyViolations := []string{}
	emptyResult := allViolations(emptyViolations)
	emptyExpectedResult := ""

	if emptyResult != nil || emptyExpectedResult != "" {
		t.Errorf("allViolations(%v) = %v; want %s", emptyViolations, emptyResult, emptyExpectedResult)
	}
}