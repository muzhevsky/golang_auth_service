package repositories

import (
	"context"
	"smartri_app/internal"
	"smartri_app/internal/infrastructure/datasources"
)

type userTestRepository struct {
	checkIfUserHasAnswersByAccountIdCommand datasources.ICheckIfUserHasAnswersByAccountIdCommand
}

func NewUserTestAnswersRepository(checkIfUserHasAnswersByAccountIdCommand datasources.ICheckIfUserHasAnswersByAccountIdCommand) internal.IUserAnswersRepository {
	return &userTestRepository{checkIfUserHasAnswersByAccountIdCommand: checkIfUserHasAnswersByAccountIdCommand}
}

func (u *userTestRepository) CheckUserHasAnswers(context context.Context, accountId int) (bool, error) {
	return u.checkIfUserHasAnswersByAccountIdCommand.Execute(context, accountId)
}
