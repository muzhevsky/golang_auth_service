package usecases

import (
	"context"
	"smartri_app/internal"
)

type checkUserPassedTestUseCase struct {
	repo internal.IUserAnswersRepository
}

func NewCheckUserHasPassedTestYetUseCase(repo internal.IUserAnswersRepository) *checkUserPassedTestUseCase {
	return &checkUserPassedTestUseCase{repo: repo}
}

func (uc *checkUserPassedTestUseCase) Check(context context.Context, accountId int) (bool, error) {
	has, err := uc.repo.CheckUserHasAnswers(context, accountId)
	if err != nil {
		return false, err
	}
	return has, nil
}
