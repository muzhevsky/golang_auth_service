package entities

import (
	"math/rand"
	"strings"
	"time"
)

const (
	verificationCodeSize  = 5
	verificationCodeRunes = "abcdefghijklmnopqrstuvwxyz0123456789"
)

type Verification struct {
	Id             int
	UserId         int
	Code           string
	ExpirationTime time.Time
}

func GenerateVerification(userId int) *Verification {
	code := make([]uint8, verificationCodeSize)
	for i := 0; i < verificationCodeSize; i++ {
		randNum := rand.Int31n(int32(len(verificationCodeRunes)))
		code[i] = verificationCodeRunes[randNum]
	}

	return &Verification{0, userId, string(code), time.Now().Add(time.Minute * time.Duration(10))}
}

func (v *Verification) ValidateVerification(verification *Verification) bool {
	return strings.ToLower(verification.Code) == v.Code
}
