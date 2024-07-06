package account

import (
	"authorization/internal/errs"
	"fmt"
)

type Nickname string

const (
	minNicknameLen = 3
	maxNicknameLen = 16
)

func (n Nickname) Validate() error {
	if !n.validateNicknameLength() {
		return fmt.Errorf("%w. nickname length can't be less than %d OR more %d", errs.ValidationError, minNicknameLen, maxNicknameLen)
	}
	if !n.validateProhibitedCharacters() {
		return fmt.Errorf("%w. nickname can contain Latin alphabet characters, numbers, underscores and hyphens", errs.ValidationError)
	}
	return nil
}

func (n Nickname) validateNicknameLength() bool {
	return runeLengthInRange([]rune(n), minNicknameLen, maxNicknameLen)
}

func (n Nickname) validateProhibitedCharacters() bool {
	for _, char := range n {
		if !(isLatinLetter(char) || isDigit(char) || isCyrillicLetter(char) || char == '-' || char == '_') {
			return false
		}
	}
	return true
}
