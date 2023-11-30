package abstraction

import "authorization/core/domain/entities"

type UserRepository interface {
	Insert(user *entities.User)
	SelectByNickname(nickname string) *entities.User
	SelectByEmail(email string) *entities.User
}
