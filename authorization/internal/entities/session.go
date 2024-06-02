package entities

import (
	"time"
)

type Session struct {
	UserId         int
	DeviceIdentity string
	AccessToken    string
	RefreshToken   string
	ExpireAt       time.Time
}
