package entities

import "github.com/google/uuid"

type User struct {
	Id                  uuid.UUID
	Login               string
	Email               string
	Password            string
	AccountCreationDate int64
}
