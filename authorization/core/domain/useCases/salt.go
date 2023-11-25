package useCases

import (
	"authorization/utils/errorHandling"
	"crypto/rand"
	"encoding/hex"
)

type SaltGenerationService interface {
	GenerateSalt() string
}

type basicSaltGenerationService struct {
}

func NewSaltService() *basicSaltGenerationService {
	return &basicSaltGenerationService{}
}

func (saltGenerationService *basicSaltGenerationService) GenerateSalt() string {
	size := 8
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	errorHandling.LogError(err)
	return hex.EncodeToString(bytes)
}
