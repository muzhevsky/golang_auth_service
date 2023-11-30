package repositories

import (
	"authorization/core/domain/entities"
	"authorization/core/external/databases/postgresql"
)

type userRepository struct {
	core *postgresql.RepositoryCore
}

func NewUserRepository(core *postgresql.RepositoryCore) *userRepository {
	return &userRepository{core}
}

func (u *userRepository) SelectByNickname(nickname string) *entities.User {
	return nil
}

func (u *userRepository) Insert(user *entities.User) {

}

func (u *userRepository) SelectByEmail(email string) *entities.User {
	return nil
}
