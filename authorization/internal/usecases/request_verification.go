package usecases

import (
	"authorization/internal"
	"authorization/internal/entities"
	"authorization/internal/errs"
	"authorization/internal/infrastructure/services/mailers"
	"context"
)

type requestVerificationUseCase struct {
	userRepo         internal.IAccountRepository
	verificationRepo internal.IVerificationRepository
	mailer           mailers.IVerificationMailer
}

func NewRequestVerificationUseCase(userRepo internal.IAccountRepository, verificationRepo internal.IVerificationRepository, mailer mailers.IVerificationMailer) *requestVerificationUseCase {
	return &requestVerificationUseCase{userRepo: userRepo, verificationRepo: verificationRepo, mailer: mailer}
}

func (u *requestVerificationUseCase) RequestVerification(context context.Context, userId int) (string, error) {
	user, err := u.userRepo.FindById(context, userId)
	if err != nil {
		return "", err
	}

	if user.IsVerified {
		return "", errs.UserIsAlreadyVerified
	}

	verification := entities.GenerateVerification(userId)
	err = u.verificationRepo.Create(context, verification)
	if err != nil {
		return "", err
	}

	u.mailer.SendMail(string(user.Email), verification.Code)

	return verification.Code, nil
	return "", nil
}
