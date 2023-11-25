package useCases

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordEncryptionService interface {
	Encrypt(string) (string, error)
}

type basicEncryptionService struct {
}

func NewEncryptionService() *basicEncryptionService {
	return &basicEncryptionService{}
}

func (service *basicEncryptionService) Encrypt(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	return string(hashedPassword), err
}
