package usecases

import (
	"errors"
	"testing"
)

// TestZero calls greetings.Hello with a name, checking
// for a valid return value.
func TestZero(t *testing.T) {
	result, err := divide(5, 0)
	t.Logf("res: %v, err: %v", result, err)
}

func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("zero division")
	}
	return a / b, nil
}
