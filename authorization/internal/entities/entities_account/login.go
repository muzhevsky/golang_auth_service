package entities_account

import (
	"authorization/internal/errs"
	"fmt"
)

type Login string

const (
	minLoginLen = 3
	maxLoginLen = 20
)

func (l Login) Validate() error {
	if !l.validateLength() {
		return fmt.Errorf("%w. login length can't be less than  %d OR more %d", errs.ValidationError, minLoginLen, maxLoginLen)
	}
	if !l.validateFirstCharacter() {
		return fmt.Errorf("%w. login must start with letter", errs.ValidationError)
	}
	if !l.validateProhibitedCharacters() {
		return fmt.Errorf("%w. login can contain Latin alphabet characters, numbers, underscores and hyphens", errs.ValidationError)
	}
	return nil
}

func (l Login) validateLength() bool {
	return stringLengthInRange(string(l), minLoginLen, maxLoginLen)
}
func (l Login) validateFirstCharacter() bool {
	return isLatinLetter(rune((l)[0]))
}
func (l Login) validateProhibitedCharacters() bool {
	for _, char := range l {
		if !(isLatinLetter(char) || isDigit(char) || char == '-' || char == '_') {
			return false
		}
	}
	return true
}
