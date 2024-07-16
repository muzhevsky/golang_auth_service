package user_data_entities

import (
	"regexp"
)

func RuneLengthInRange(str []rune, min, max int) bool {
	return len(str) >= min && len(str) <= max
}
func stringLengthInRange(str string, min, max int) bool {
	return RuneLengthInRange([]rune(str), min, max)
}
func IsLatinLetter(symbol rune) bool {
	return (symbol >= 'a' && symbol <= 'z') || (symbol >= 'A' && symbol <= 'Z')
}
func IsCyrillicLetter(char rune) bool {
	return (char >= 'а' && char <= 'я') || (char >= 'А' && char <= 'Я')
}
func IsDigit(symbol rune) bool {
	return symbol >= '0' && symbol <= '9'
}
func isEmail(str string) bool {
	m, err := regexp.Match("^[а-яА-Яa-zA-Z0-9.+-]+@([a-zа-я-]+\\.)+[a-zа-я-]{2,4}$", []byte(str))
	if err != nil {
		return false
	}
	return m
}
