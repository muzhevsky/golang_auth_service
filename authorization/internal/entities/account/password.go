package account

import (
	"authorization/internal/errs"
	"fmt"
)

type Password string

const (
	minPasswordLen = 8
	maxPasswordLen = 20
)

func (p *Password) Validate() error {
	if !p.validateLength() {
		return fmt.Errorf("%w, password length can't be less than %d OR more %d", errs.PasswordValidationError, minPasswordLen, maxPasswordLen)
	}
	if !p.validateCharacters() {
		return fmt.Errorf("%w. password should contain at least one Latin alphabet character and one digit", errs.PasswordValidationError)
	}
	return nil
}

func (p *Password) validateLength() bool {
	return stringLengthInRange(string(*p), minPasswordLen, maxPasswordLen)
}

func (p *Password) validateCharacters() bool {
	digitCount, letterCount := 0, 0
	for _, char := range string(*p) {
		if isLatinLetter(char) {
			letterCount++
		} else if isDigit(char) {
			digitCount++
		}
	}
	if digitCount > 0 && letterCount > 0 {
		return true
	}
	return false
}
