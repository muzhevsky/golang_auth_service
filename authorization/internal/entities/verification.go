package entities

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	verificationCodeSize  = 5
	verificationCodeRunes = "abcdefghijklmnopqrstuvwxyz0123456789"
)

type Verification struct {
	UserId      int
	Code        string
	ExpiredTime time.Time
}

func GenerateVerification(userId int, expiredTime time.Time) *Verification {
	code := make([]uint8, verificationCodeSize)
	for i := 0; i < verificationCodeSize; i++ {
		randNum := rand.Int31n(int32(len(verificationCodeRunes)))
		code[i] = verificationCodeRunes[randNum]
	}

	return &Verification{userId, string(code), expiredTime}
}

func (v *Verification) VerifyUser(verification *Verification) bool {
	return strings.ToLower(verification.Code) == v.Code
}

func (v *Verification) GenerateEmailBody() string {
	return fmt.Sprintf("Код верификации: %s", v.Code)
}
