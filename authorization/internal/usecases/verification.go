package usecases

import (
	"authorization/internal/entities"
	"context"
	"time"
)

type verificationUseCase struct {
	userRepo         IUserRepo
	verificationRepo IVerificationRepo
	mailer           IMailer
}

func NewVerificationUseCase(userRepo IUserRepo, verificationRepo IVerificationRepo, mailer IMailer) IVerification {
	return &verificationUseCase{userRepo, verificationRepo, mailer}
}

// CreateVerification - creates new verification record in repository
// returns repository error
func (v *verificationUseCase) CreateVerification(context context.Context, user *entities.User) error {
	err := v.verificationRepo.Clear(context, user.Id)
	if err != nil {
		return err
	}

	verification := entities.GenerateVerification(user.Id, time.Now().Add(time.Minute*time.Duration(10)))
	err = v.verificationRepo.Create(context, verification)
	if err != nil {
		return err
	}

	v.mailer.SendMail(user.EMail, "Подтверждение учетной записи", verification.GenerateEmailBody())
	return nil
}

// Verify - serves verification process, checking if there's any verification records in repository by provided userId
// within verification object.
// returns:
//   - ExpiredCode error if there are no active verification codes
//   - WrongVerificationCode error if code is wrong
func (v *verificationUseCase) Verify(context context.Context, verification *entities.Verification) error {
	existingVerification, err := v.verificationRepo.FindOne(context, verification.UserId)
	if err != nil {
		return err
	}

	result := existingVerification.VerifyUser(verification)
	if !result {
		return entities.WrongVerificationCode
	}

	if existingVerification.ExpiredTime.Before(time.Now()) {
		return entities.ExpiredCode
	}
	err = v.userRepo.Verify(context, verification.UserId)
	if err != nil {
		return err
	}
	return nil
}
