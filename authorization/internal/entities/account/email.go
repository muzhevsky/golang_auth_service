package account

import (
	"authorization/internal/errs"
	"fmt"
)

type Email string

const (
	minEmailLen = 6
	maxEmailLen = 254
)

func (e *Email) Validate() error {
	if !e.validateLength() {
		return fmt.Errorf("%w email length can't be less than  %d OR more %d", errs.EmailValidationError, minEmailLen, maxEmailLen)
	}
	if !e.validateFormat() {
		return fmt.Errorf("%w wrong email format", errs.EmailValidationError)
	}
	return nil
}

func (e *Email) validateLength() bool {
	return stringLengthInRange(string(*e), minEmailLen, maxEmailLen)
}

func (e *Email) validateFormat() bool {
	return isEmail(string(*e))
}
