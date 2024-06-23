package repositories

import (
	"authorization/internal/entities"
	"authorization/internal/infrastructure/datasources"
	"context"
)

type userRepo struct {
	accountDS datasources.IAccountDataSource
}

func NewUserRepo(accountDS datasources.IAccountDataSource) *userRepo {
	return &userRepo{accountDS}
}

func (u *userRepo) Create(context context.Context, user *entities.Account) (id int, err error) {
	return u.accountDS.Create(context, user)
}

func (u *userRepo) FindById(context context.Context, id int) (*entities.Account, error) {
	return u.accountDS.SelectById(context, id)
}

func (u *userRepo) FindByLogin(context context.Context, login string) (*entities.Account, error) {
	return u.accountDS.SelectByLogin(context, login)
}

func (u *userRepo) FindByEmail(context context.Context, email string) (*entities.Account, error) {
	return u.accountDS.SelectByEmail(context, email)
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

func (u *userRepo) Update(context context.Context, account *entities.Account) error {
	return u.accountDS.UpdateById(context, account.Id, func(u *entities.Account) {
		u.Login = account.Login
		u.EMail = account.EMail
		u.Password = account.Password
		u.IsVerified = account.IsVerified
		u.Nickname = account.Nickname
	})
}
