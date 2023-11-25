package entities

import (
	"authorization/internal/domain/errors"
	"authorization/internal/domain/utils"
)

type charactersValidator struct {
	allowedChars int
}

func (v *charactersValidator) IsValid(login string) (bool, error) {
	lettersAllowed := (1<<utils.Letter)&v.allowedChars > 0
	digitsAllowed := (1<<utils.Digit)&v.allowedChars > 0
	specialCharsAllowed := (1<<utils.SpecialCharacter)&v.allowedChars > 0

	checkers := make([]func(rune) bool, 3)
	if lettersAllowed && digitsAllowed && specialCharsAllowed {
		checkers = append(checkers, func(r rune) bool {
			return true
		})
	} else {
		if lettersAllowed {
			checkers = append(checkers, utils.RuneIsLetter)
		}
		if digitsAllowed {
			checkers = append(checkers, utils.RuneIsDigit)
		}
		if specialCharsAllowed {
			checkers = append(checkers, utils.RuneIsSpecial)
		}
	}

	for i := 0; i < len(login); i++ {
		for k := 0; k < len(checkers); k++ {
			if !checkers[k](rune(login[i])) {
				return false, errors.LoginContentError
			}
		}
	}

	return true, nil
}
