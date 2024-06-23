package entities

import (
	"time"
)

type Account struct {
	Id           int
	Login        string
	Password     string
	EMail        string
	Nickname     string
	IsVerified   bool
	CreationTime time.Time
}

func (u *Account) Verify() {
	u.IsVerified = true
}
