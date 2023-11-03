package validator

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}

func ValidateInt(value, minValue, maxValue int) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("must be in range from %d-%d", minValue, maxValue)
	}
	return nil
}
func ValidateInt32(value, minValue, maxValue int32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("must be in range from %d-%d", minValue, maxValue)
	}
	return nil
}

func ValidateInt64(value, minValue, maxValue int64) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("must be in range from %d-%d", minValue, maxValue)
	}

	len := len(strconv.Itoa(int(value)))
	if len < 10 {
		return fmt.Errorf("must be in range from %d-%d", minValue, maxValue)
	}
	return nil
}

func ValidateDouble(value, minValue float64, precision int) error {
	if value < minValue {
		return fmt.Errorf("must be in greater than %.2f", minValue)
	}

	str := strconv.FormatFloat(value, 'f', -1, 64)
	parts := strings.Split(str, ".")
	fractionalDigits := len(parts[1])

	if fractionalDigits > 2 {
		return fmt.Errorf("value has a precision greater than %d", precision)
	}
	return nil
}

