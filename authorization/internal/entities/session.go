package entities

import (
	"time"
)

type Session struct {
	AccessToken  string
	RefreshToken string
	ExpireAt     time.Time
}
