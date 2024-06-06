package entities

import (
	"time"
)

type Session struct {
	Id             int
	UserId         int
	DeviceIdentity string
	AccessToken    string
	RefreshToken   string
	ExpiresAt      time.Time
}
