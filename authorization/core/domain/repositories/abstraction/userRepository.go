package abstraction

import "authorization/core/domain/entities"

type UserRepository interface {
	Insert(user *entities.User) error
	SelectByLogin(nickname string) (*entities.User, error)
	SelectByEmail(email string) (*entities.User, error)
}
