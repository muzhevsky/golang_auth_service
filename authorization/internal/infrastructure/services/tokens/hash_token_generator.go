package tokens

import (
	"math/big"
	"math/rand"
)

type hashTokenGenerator struct {
	hashProvider IHashProvider
}

func NewHashRefreshTokenGenerator(hashProvider IHashProvider) *hashTokenGenerator {
	return &hashTokenGenerator{hashProvider: hashProvider}
}

func (h *hashTokenGenerator) GenerateToken(accountId int) (string, error) {
	var number *big.Int
	randomNumber := big.NewInt(rand.Int63())
	number = big.NewInt(int64(accountId))
	number.Lsh(number, 64)
	number.Add(number, randomNumber)
	stringToHash := number.String()
	hash, err := h.hashProvider.GenerateHash(stringToHash)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
