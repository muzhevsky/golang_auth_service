package repositories

import (
	"context"
	"smartri_app/internal/entities"
	"smartri_app/internal/infrastructure/datasources"
)

type userRepository struct {
	ds datasources.IUserDataSource
}

func NewUserRepository(ds datasources.IUserDataSource) *userRepository {
	return &userRepository{ds: ds}
}

func (u *userRepository) GetByAccountId(context context.Context, accountId int) (*entities.User, error) {
	return u.ds.SelectByAccountId(context, accountId)
}

func (u *userRepository) AddOrUpdate(context context.Context, details *entities.User) error {
	data, err := u.GetByAccountId(context, details.AccountId)
	if err != nil {
		return err
	}

	if data == nil {
		err := u.ds.Insert(context, details)
		if err != nil {
			return err
		}
	} else {
		details.AccountId = data.AccountId
		details.XP = data.XP
		err := u.ds.UpdateByAccountId(context, details.AccountId, details)
		if err != nil {
			return err
		}
	}

	return nil
}
