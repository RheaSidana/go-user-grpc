package users

import (
	"errors"
	"fmt"
)

func fieldViolation(field string, err error) (string) {
	return fmt.Sprintf("{ Field: %s, \n Description: %s, \n}\n", field, err.Error())
}

func allViolations(violations []string) (error) {
	if len(violations) == 0 {
		return nil
	}

	var errs string
	for _, violation := range violations{
		errs += violation
	}

	return errors.New(errs)
}
