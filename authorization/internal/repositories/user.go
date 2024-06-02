package repositories

import (
	"authorization/internal/entities"
	"authorization/internal/infrastructure/datasources"
	"context"
)

type userRepo struct {
	userDS datasources.IUserDataSource
}

func NewUserRepo(userDS datasources.IUserDataSource) *userRepo {
	return &userRepo{userDS}
}

func (u *userRepo) Create(context context.Context, user *entities.User) (id int, err error) {
	return u.userDS.Create(context, user)
}

func (u *userRepo) FindById(context context.Context, id int) (*entities.User, error) {
	return u.userDS.SelectById(context, id)
}

func (u *userRepo) FindByLogin(context context.Context, login string) (*entities.User, error) {
	return u.userDS.SelectByLogin(context, login)
}

func (u *userRepo) FindByEmail(context context.Context, email string) (*entities.User, error) {
	return u.userDS.SelectByEmail(context, email)
}

func (u *userRepo) CheckLoginExist(context context.Context, login string) (result bool, err error) {
	user, err := u.FindByLogin(context, login)

	if err != nil {
		return false, err
	}

	return user != nil, nil
}

func (u *userRepo) CheckEmailExist(context context.Context, email string) (bool, error) {
	user, err := u.FindByEmail(context, email)

	if err != nil {
		return false, err
	}

	return user != nil, nil
}

func (u *userRepo) Verify(context context.Context, id int) error {
	err := u.userDS.UpdateById(context, id, func(user *entities.User) {
		user.IsVerified = true
	})

	return err
}
