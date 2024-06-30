package usecases

import (
	"authorization/internal"
	"context"
)

type checkVerificationUsecase struct {
	accountRepository internal.IAccountRepository
}

func NewCheckVerificationUsecase(accountRepository internal.IAccountRepository) *checkVerificationUsecase {
	return &checkVerificationUsecase{accountRepository: accountRepository}
}

func (uc *checkVerificationUsecase) Check(context context.Context, accountId int) (bool, error) {
	account, err := uc.accountRepository.FindById(context, accountId)
	if err != nil {
		return false, err
	}

	return account.IsVerified, nil
}
