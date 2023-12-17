package usecase

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

func (v *verificationUseCase) CreateVerification(user *entities.User) error {
	verification := entities.GenerateVerification(user.Id, time.Now().Add(time.Minute*time.Duration(10)))
	err := v.verificationRepo.Create(verification)
	if err != nil {
		return err
	}

	v.mailer.SendMail(user.EMail, "Подтверждение учетной записи", verification.GenerateEmailBody())
	return nil
}

func (v *verificationUseCase) Verify(context context.Context, verification *entities.Verification) (result bool, err error) {
	existingVerification, err := v.verificationRepo.FindOne(verification.UserId)
	if err != nil {
		return false, err
	}

	if existingVerification.ExpiredTime.Before(time.Now()) {
		return false, entities.ExpiredCode
	}

	result = existingVerification.VerifyUser(verification)
	if result {
		err = v.userRepo.Verify(context, verification.UserId)
		if err != nil {
			return false, err
		}
	}

	return result, nil
}
