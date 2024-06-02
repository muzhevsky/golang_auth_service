package entities

import (
	"authorization/internal/errs"
	"fmt"
	"net/mail"
)

const (
	minLoginLen    = 3
	maxLoginLen    = 20
	minPasswordLen = 8
	maxPasswordLen = 20
	minEmailLen    = 6
	maxEmailLen    = 254
	minNicknameLen = 3
	maxNicknameLen = 16
)

type UserValidator struct {
}

func (v *UserValidator) ValidateLogin(login string) error {
	if !v.validateLoginLength(login) {
		return fmt.Errorf("%w. login length can't be less than  %d OR more %d", errs.ValidationError, minLoginLen, maxLoginLen)
	}
	if !v.validateFirstCharacter(rune(login[0])) {
		return fmt.Errorf("%w. login must start with letter", errs.ValidationError)
	}
	if !v.validateProhibitedCharacters(login) {
		return fmt.Errorf("%w. login can contain Latin alphabet characters, numbers, underscores and hyphens", errs.ValidationError)
	}
	return nil
}

func (v *UserValidator) validateLoginLength(login string) bool {
	return v.stringLengthInRange(login, minLoginLen, maxLoginLen)
}

func (v *UserValidator) validateFirstCharacter(char rune) bool {
	return v.isLetter(char)
}

func (v *UserValidator) validateProhibitedCharacters(source string) bool {
	for _, char := range source {
		if !(v.isLetter(char) || v.isDigit(char) || char == '-' || char == '_') {
			return false
		}
	}
	return true
}

func (v *UserValidator) ValidateEmail(email string) error {
	if !v.validateLengthEmail(email) {
		return fmt.Errorf("%w email length can't be less than  %d OR more %d", errs.ValidationError, minEmailLen, maxEmailLen)
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("%w %s", errs.ValidationError, err)
	}
	return nil
}

func (v *UserValidator) validateLengthEmail(email string) bool {
	return v.stringLengthInRange(email, minEmailLen, maxEmailLen)
}

func (v *UserValidator) ValidatePassword(password string) error {
	if !v.validatePasswordLength(password) {
		return fmt.Errorf("%w, password length can't be less than  %d OR more %d", errs.ValidationError, minPasswordLen, maxPasswordLen)
	}
	if !v.validatePasswordCharacters(password) {
		return fmt.Errorf("%w. password should contain at least one Latin alphabet character and one number", errs.ValidationError) // todo задать нормальный exception
	}
	return nil
}

func (v *UserValidator) validatePasswordLength(password string) bool {
	return v.stringLengthInRange(password, minPasswordLen, maxPasswordLen)
}

func (v *UserValidator) validatePasswordCharacters(password string) bool {
	digitCount, letterCount := 0, 0
	for _, char := range password {
		if v.isLetter(char) {
			letterCount++
		} else if v.isDigit(char) {
			digitCount++
		} else {
			return false
		}
	}
	if digitCount > 0 && letterCount > 0 {
		return true
	}
	return false
}
func (v *UserValidator) ValidateNickname(nickname string) error {
	if !v.validateNicknameLength(nickname) {
		return fmt.Errorf("%w. nickname length can't be less than  %d OR more %d", errs.ValidationError, minNicknameLen, maxNicknameLen)
	}
	if !v.validateProhibitedCharacters(nickname) { // todo другие правила валидации
		return fmt.Errorf("%w. nickname can contain Latin alphabet characters, numbers, underscores and hyphens", errs.ValidationError)
	}
	return nil
}

func (v *UserValidator) validateNicknameLength(nickname string) bool {
	return v.stringLengthInRange(nickname, minNicknameLen, maxNicknameLen)
}

func (v *UserValidator) stringLengthInRange(str string, min, max int) bool {
	return len(str) >= min && len(str) <= max
}
func (v *UserValidator) isLetter(symbol rune) bool {
	return (symbol >= 'a' && symbol <= 'z') || (symbol >= 'A' && symbol <= 'Z')
}

func (v *UserValidator) isDigit(symbol rune) bool {
	return symbol >= '0' && symbol <= '9'
}
