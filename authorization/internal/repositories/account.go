package repositories

import (
	"authorization/internal"
	"authorization/internal/entities/account"
	"authorization/internal/infrastructure/datasources"
	"context"
)

type accountRepo struct {
	selectAccountByIdCommand    datasources.ISelectAccountByIdCommand
	selectAccountByEmailCommand datasources.ISelectAccountByEmailCommand
	selectAccountByLoginCommand datasources.ISelectAccountByLoginCommand
	updateAccountByIdCommand    datasources.IUpdateAccountByIdCommand
	insertAccountCommand        datasources.IInsertAccountCommand
}

func NewAccountRepository(
	selectAccountByIdCommand datasources.ISelectAccountByIdCommand,
	selectAccountByEmailCommand datasources.ISelectAccountByEmailCommand,
	selectAccountByLoginCommand datasources.ISelectAccountByLoginCommand,
	updateAccountByIdCommand datasources.IUpdateAccountByIdCommand,
	insertAccountCommand datasources.IInsertAccountCommand) internal.IAccountRepository {
	return &accountRepo{
		selectAccountByIdCommand:    selectAccountByIdCommand,
		selectAccountByEmailCommand: selectAccountByEmailCommand,
		selectAccountByLoginCommand: selectAccountByLoginCommand,
		updateAccountByIdCommand:    updateAccountByIdCommand,
		insertAccountCommand:        insertAccountCommand}
}

func (u *accountRepo) Create(context context.Context, user *account.Account) (id int, err error) {
	return u.insertAccountCommand.Execute(context, user)
}

func (u *accountRepo) FindById(context context.Context, id int) (*account.Account, error) {
	return u.selectAccountByIdCommand.Execute(context, id)
}

func (u *accountRepo) FindByLogin(context context.Context, login account.Login) (*account.Account, error) {
	return u.selectAccountByLoginCommand.Execute(context, string(login))
}

func (u *accountRepo) FindByEmail(context context.Context, email account.Email) (*account.Account, error) {
	return u.selectAccountByEmailCommand.Execute(context, string(email))
}

func (u *accountRepo) CheckLoginExist(context context.Context, login account.Login) (result bool, err error) {
	user, err := u.FindByLogin(context, login)

	if err != nil {
		return false, err
	}

	return user != nil, nil
}

func (u *accountRepo) CheckEmailExist(context context.Context, email account.Email) (bool, error) {
	user, err := u.FindByEmail(context, email)

	if err != nil {
		return false, err
	}

	return user != nil, nil
}

func (u *accountRepo) UpdateById(context context.Context, id int, account *account.Account) error {
	return u.updateAccountByIdCommand.Execute(context, id, account)
}
