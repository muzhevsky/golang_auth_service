package entities

import (
	"time"
)

type Session struct {
	Id             int
	AccountId      int
	DeviceIdentity string
	AccessToken    string
	RefreshToken   string
	ExpiresAt      time.Time
}
