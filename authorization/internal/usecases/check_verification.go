package usecases

import (
	"authorization/internal"
	"authorization/internal/errs"
	"context"
)

type checkVerificationUsecase struct {
	accountRepository internal.IAccountRepository
}

func NewCheckVerificationUsecase(accountRepository internal.IAccountRepository) internal.ICheckVerificationUseCase {
	return &checkVerificationUsecase{accountRepository: accountRepository}
}

func (uc *checkVerificationUsecase) CheckVerification(context context.Context, accountId int) (bool, error) {
	account, err := uc.accountRepository.FindById(context, accountId)
	if err != nil {
		return false, err
	}
	if account == nil {
		return false, errs.AccountNotFound
	}

	return account.IsVerified, nil
}
