package repositories

import (
	"context"
	"smartri_app/internal"
	"smartri_app/internal/entities/user_data"
	"smartri_app/internal/infrastructure/datasources"
)

type userDataRepository struct {
	selectUserDataByAccountIdCommand datasources.ISelectUserDataByAccountIdCommand
	updateUserDataCommand            datasources.IUpdateUserDataByAccountIdCommand
	insertUserDataCommand            datasources.IInsertUserDataCommand
}

func NewUserDataRepository(selectUserDataByAccountIdCommand datasources.ISelectUserDataByAccountIdCommand, updateUserDataCommand datasources.IUpdateUserDataByAccountIdCommand, insertUserDataCommand datasources.IInsertUserDataCommand) internal.IUserDataRepository {
	return &userDataRepository{
		selectUserDataByAccountIdCommand: selectUserDataByAccountIdCommand,
		updateUserDataCommand:            updateUserDataCommand,
		insertUserDataCommand:            insertUserDataCommand}
}

func (u *userDataRepository) GetByAccountId(context context.Context, accountId int) (*user_data.UserData, error) {
	return u.selectUserDataByAccountIdCommand.Execute(context, accountId)
}

func (u *userDataRepository) Update(context context.Context, details *user_data.UserData) (*user_data.UserData, error) {
	updated, err := u.updateUserDataCommand.Execute(context, details)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (u *userDataRepository) Create(context context.Context, details *user_data.UserData) error {
	err := u.insertUserDataCommand.Execute(context, details)
	return err
}
