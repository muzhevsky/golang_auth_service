package infrastructure

import (
	"authorization/internal/usecases"
	"math/big"
	"math/rand"
)

type hashRefreshTokenGenerator struct {
	hashProvider usecases.IHashProvider
}

func NewHashRefreshTokenGenerator(hashProvider usecases.IHashProvider) *hashRefreshTokenGenerator {
	tokenGenerator := &hashRefreshTokenGenerator{hashProvider: hashProvider}
	return tokenGenerator
}

func (h *hashRefreshTokenGenerator) GenerateToken(userId int) (string, error) {
	var number *big.Int
	randomNumber := big.NewInt(rand.Int63())
	number = big.NewInt(int64(userId))
	number.Lsh(number, 64)
	number.Add(number, randomNumber)
	stringToHash := number.String()
	hash, err := h.hashProvider.GenerateHash(stringToHash)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
