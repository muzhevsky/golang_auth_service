package infrastructure

import "golang.org/x/crypto/bcrypt"

type bcryptHashProvider struct {
	hashCost int
}

func NewBcryptHashProvider(hashCost int) *bcryptHashProvider {
	return &bcryptHashProvider{hashCost: hashCost}
}

func (p *bcryptHashProvider) GenerateHashPassword(stringToHash string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(stringToHash), p.hashCost)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (p *bcryptHashProvider) CompareStringAndHash(stringToCompare string, hashedString string) bool {

	passwordMatched := bcrypt.CompareHashAndPassword([]byte(hashedString), []byte(stringToCompare))
	if passwordMatched != nil {
		return false
	}

	return true
}
