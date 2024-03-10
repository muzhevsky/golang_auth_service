package entities

import (
	"time"
)

type Session struct {
	UserId            int
	DeviceDescription string
	AccessToken       string
	RefreshToken      string
	ExpireAt          time.Time
}
