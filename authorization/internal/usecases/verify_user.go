package usecases

import (
	"authorization/internal"

	verificationpkg "authorization/internal/entities/verification"
	"authorization/internal/errs"
	"authorization/internal/infrastructure/services/mailers"
	"context"
	"time"
)

type verificationUseCase struct {
	userRepo         internal.IAccountRepository
	verificationRepo internal.IVerificationRepository
	mailer           mailers.IVerificationMailer
}

func NewVerificationUseCase(
	userRepo internal.IAccountRepository,
	verificationRepo internal.IVerificationRepository,
	mailer mailers.IVerificationMailer) internal.IVerifyUserUseCase {
	return &verificationUseCase{userRepo, verificationRepo, mailer}
}

// Verify - serves verification process, checking if there's any verification records in repository by provided userId
// within verification object.
// returns:
//   - ExpiredVerificationCode error if there are no active verification codes
//   - WrongVerificationCode error if code is wrong
func (v *verificationUseCase) Verify(context context.Context, accountId int, code string) error {
	verification := &verificationpkg.Verification{
		AccountId: accountId,
		Code:      code,
	}

	existingVerifications, err := v.verificationRepo.FindByAccountId(context, verification.AccountId)
	if err != nil {
		return err
	}

	var existingVerification *verificationpkg.Verification
	for i := 0; i < len(existingVerifications); i++ {
		if existingVerifications[i].Code == code {
			existingVerification = existingVerifications[i]
			break
		}
	}

	if existingVerification == nil {
		return errs.WrongVerificationCode
	}

	if existingVerification.ExpirationTime.Before(time.Now()) {
		return errs.ExpiredVerificationCode
	}

	result := existingVerification.ValidateVerification(verification)
	if !result {
		return errs.WrongVerificationCode
	}

	account, err := v.userRepo.FindById(context, accountId)
	if err != nil {
		return err
	}

	account.Verify()

	err = v.userRepo.UpdateById(context, accountId, account)
	if err != nil {
		return err
	}

	err = v.verificationRepo.Clear(context, accountId)
	if err != nil {
		return err
	}

	return nil
}
