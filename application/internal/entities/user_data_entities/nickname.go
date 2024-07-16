package user_data_entities

import (
	"fmt"
	"smartri_app/internal/errs"
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
	return RuneLengthInRange([]rune(n), minNicknameLen, maxNicknameLen)
}

func (n Nickname) validateProhibitedCharacters() bool {
	for _, char := range n {
		if !(IsLatinLetter(char) || IsDigit(char) || IsCyrillicLetter(char) || char == '-' || char == '_') {
			return false
		}
	}
	return true
}
