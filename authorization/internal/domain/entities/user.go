package entities

import "time"

type User struct {
	Login                   string `json:"login"`
	Email                   string `json:"email"`
	Password                string `json:"password"`
	AccountCreationDateTime int64
}

func CreateUser(login string, email string, password string) *User {
	return &User{login, email, password, time.Now().Unix()}
}
