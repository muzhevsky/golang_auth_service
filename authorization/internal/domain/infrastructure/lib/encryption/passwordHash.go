package encryption

import "golang.org/x/crypto/bcrypt"

type bcryptHashGenerator struct {
	bcryptCost int
}

func New(bcryptCost int) *bcryptHashGenerator {
	return &bcryptHashGenerator{bcryptCost}
}

func (generator *bcryptHashGenerator) EncryptString(str string) (string, error) {
	hashedString, err := bcrypt.GenerateFromPassword([]byte(str), generator.bcryptCost)
	return string(hashedString), err
}
