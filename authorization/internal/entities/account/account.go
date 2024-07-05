package account

import (
	"time"
)

type Account struct {
	Id               int
	Login            *Login
	Password         *Password
	Email            *Email
	Nickname         *Nickname
	IsVerified       bool
	RegistrationDate time.Time
}

func (u *Account) Verify() {
	u.IsVerified = true
}

func (u *Account) Validate() error {
	err := u.Login.Validate()
	if err != nil {
		return err
	}
	err = u.Email.Validate()
	if err != nil {
		return err
	}
	err = u.Password.Validate()
	if err != nil {
		return err
	}
	err = u.Nickname.Validate()
	if err != nil {
		return err
	}
	return nil
}
