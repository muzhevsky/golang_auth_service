package useCases

import (
	"authorization/internal/domain/entities"
)

type (
	UserRepository interface {
		Insert(user *entities.User) error
		SelectByLogin(nickname string) (*entities.User, error)
		SelectByEmail(email string) (*entities.User, error)
	}

	AuthenticationUseCase interface {
		SignUp(dto *entities.User) error
	}

	StringEncryptor interface {
		EncryptString(string) (string, error)
	}
)
