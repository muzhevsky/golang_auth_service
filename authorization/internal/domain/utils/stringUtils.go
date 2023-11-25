package utils

const (
	Letter = iota
	Digit
	SpecialCharacter
)

func RuneIsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func RuneIsLetter(r rune) bool {
	lowerCase := r >= 'a' && r <= 'z'
	if lowerCase {
		return true
	}
	upperCase := r >= 'A' && r <= 'Z'
	if upperCase {
		return true
	}
	return false
}

func RuneIsSpecial(r rune) bool {
	return !RuneIsDigit(r) && !RuneIsLetter(r)
}
