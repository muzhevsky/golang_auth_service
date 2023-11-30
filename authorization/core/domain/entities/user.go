package entities

import "github.com/google/uuid"

type User struct {
	id                  uuid.UUID
	login               string
	email               string
	password            string
	salt                string
	accountCreationDate int64
}
