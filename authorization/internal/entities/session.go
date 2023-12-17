package entities

import (
	"fmt"
	"math/rand"
	"time"
)

type Session struct {
	Id           int
	AccessToken  string
	RefreshToken string
	ExpireAt     time.Time
}

func (a *Session) GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)
	src := rand.NewSource(time.Now().Unix())
	r := rand.New(src)

	_, err := r.Read(b)
	if err != nil {
		return "", fmt.Errorf("GenerateRefreshToken failed %w", err)
	}
	return fmt.Sprintf("%x", b), nil
}
