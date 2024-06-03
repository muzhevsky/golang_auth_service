package usecases

import (
	"authorization/internal"
	"authorization/internal/controllers/requests"
	"authorization/internal/entities"
	"authorization/internal/errs"
	"authorization/internal/infrastructure/services/mailers"
	"context"
	"fmt"
	"time"
)

type verificationUseCase struct {
	userRepo         internal.IUserRepository
	verificationRepo internal.IVerificationRepository
	mailer           mailers.IVerificationMailer
}

func NewVerificationUseCase(userRepo internal.IUserRepository, verificationRepo internal.IVerificationRepository, mailer mailers.IVerificationMailer) internal.IVerifyUserUseCase {
	return &verificationUseCase{userRepo, verificationRepo, mailer}
}

// Verify - serves verification process, checking if there's any verification records in repository by provided userId
// within verification object.
// returns:
//   - ExpiredCode error if there are no active verification codes
//   - WrongVerificationCode error if code is wrong
func (v *verificationUseCase) Verify(context context.Context, request *requests.VerificationRequest) error {
	verification := &entities.Verification{
		UserId: request.UserId,
		Code:   request.Code,
	}

	existingVerifications, err := v.verificationRepo.FindByUserId(context, verification.UserId)
	if err != nil {
		return err
	}

	var existingVerification *entities.Verification
	for _, verification := range existingVerifications {
		fmt.Println(verification.Code)
		if verification.Code == request.Code {
			existingVerification = verification
		}
	}

	if existingVerification == nil {
		return errs.WrongVerificationCode
	}

	if existingVerification.ExpirationTime.Before(time.Now()) {
		return errs.ExpiredCode
	}

	result := existingVerification.ValidateVerification(verification)
	if !result {
		return errs.WrongVerificationCode
	}

	err = v.userRepo.Verify(context, verification.UserId) // todo перенести метод в entity, а в репозитории сделать Update()
	if err != nil {
		return err
	}

	return nil
}
