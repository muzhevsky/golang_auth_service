package usecases

import (
	"authorization/internal"
	"authorization/internal/errs"
	"context"
)

type checkVerificationUseCase struct {
	accountRepository internal.IAccountRepository
}

func NewCheckVerificationUseCase(accountRepository internal.IAccountRepository) internal.ICheckVerificationUseCase {
	return &checkVerificationUseCase{accountRepository: accountRepository}
}

func (uc *checkVerificationUseCase) CheckVerification(context context.Context, accountId int) (bool, error) {
	account, err := uc.accountRepository.FindById(context, accountId)
	if err != nil {
		return false, err
	}
	if account == nil {
		return false, errs.AccountNotFound
	}

	return account.IsVerified, nil
}
