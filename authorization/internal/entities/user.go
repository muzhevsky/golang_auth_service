package entities

import (
	"time"
)

type User struct {
	Id           int
	Login        string
	Password     string
	EMail        string
	Nickname     string
	IsVerified   bool
	CreationTime time.Time
}
