package account

import (
	"time"
)

type Account struct {
	Id               int
	Login            Login
	Password         Password
	Email            Email
	IsVerified       bool
	RegistrationDate time.Time
}

func (a *Account) Verify() {
	a.IsVerified = true
}

func (a *Account) ConfirmCreation(hashedPassword Password) {
	a.Password = hashedPassword
	a.RegistrationDate = time.Now()
}

func (a *Account) ConfirmCreationString(hashedPassword string) {
	password := Password(hashedPassword)
	a.ConfirmCreation(password)
}

func (a *Account) ConfirmCreationBytes(hashedPassword []byte) {
	password := Password(hashedPassword)
	a.ConfirmCreation(password)
}

func (a *Account) Validate() error {
	err := a.Login.Validate()
	if err != nil {
		return err
	}
	err = a.Email.Validate()
	if err != nil {
		return err
	}
	err = a.Password.Validate()
	if err != nil {
		return err
	}
	return nil
}
