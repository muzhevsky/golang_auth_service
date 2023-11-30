package useCases

import (
	"authorization/utils/errorsAndPanics"
	"crypto/rand"
	"encoding/hex"
)

func GenerateSalt(password string) string {
	size := 8
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	errorsAndPanics.HandleError(err)
	return hex.EncodeToString(bytes)
}
