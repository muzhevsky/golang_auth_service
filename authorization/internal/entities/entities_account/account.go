package entities_account

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

// NewAccount creates new entities_account and checks if it's valid
//
// Returns:
//
// Account - pointer to the new entities_account,
//
// errs.ValidationError - if some argument is not valid
func NewAccount(login string, email string, password string) (*Account, error) {
	var result Account

	result.Login = Login(login)
	result.Email = Email(email)
	result.Password = Password(password)

	err := result.Validate()
	if err != nil {
		return nil, err
	}

	return &result, nil
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
