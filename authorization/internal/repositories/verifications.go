package repositories

import (
	"authorization/internal"
	"authorization/internal/entities/verification_entities"
	"authorization/internal/infrastructure/datasources"
	"context"
)

type verificationRepo struct {
	create            datasources.ICreateVerificationCommand
	selectByAccountId datasources.ISelectVerificationsByAccountIdCommand
	deleteByAccountId datasources.IDeleteVerificationsByAccountIdCommand
}

func NewVerificationRepo(
	createVerificationCommand datasources.ICreateVerificationCommand,
	selectVerificationsByAccountIdCommand datasources.ISelectVerificationsByAccountIdCommand,
	deleteVerificationsByAccountIdCommand datasources.IDeleteVerificationsByAccountIdCommand) internal.IVerificationRepository {
	return &verificationRepo{
		createVerificationCommand,
		selectVerificationsByAccountIdCommand,
		deleteVerificationsByAccountIdCommand}
}

func (repo *verificationRepo) Create(context context.Context, verification *verification_entities.Verification) error {
	return repo.create.Execute(context, verification)
}

func (repo *verificationRepo) FindByAccountId(context context.Context, userId int) ([]*verification_entities.Verification, error) {
	return repo.selectByAccountId.Execute(context, userId)
}

func (repo *verificationRepo) Clear(context context.Context, accountId int) error {
	return repo.deleteByAccountId.Execute(context, accountId)
}
