package usecases

import (
	"authorization/internal"
	"authorization/internal/entities"
	"authorization/internal/errs"
	"authorization/internal/infrastructure/services/mailers"
	"context"
)

type requestVerificationRequest struct {
	userRepo         internal.IAccountRepository
	verificationRepo internal.IVerificationRepository
	mailer           mailers.IVerificationMailer
}

func NewRequestVerificationRequest(userRepo internal.IAccountRepository, verificationRepo internal.IVerificationRepository, mailer mailers.IVerificationMailer) *requestVerificationRequest {
	return &requestVerificationRequest{userRepo: userRepo, verificationRepo: verificationRepo, mailer: mailer}
}

func (u *requestVerificationRequest) RequestVerification(context context.Context, userId int) (string, error) {
	user, err := u.userRepo.FindById(context, userId)
	if err != nil {
		return "", err
	}

	if user.IsVerified {
		return "", errs.UserIsAlreadyVerified
	}

	verification := entities.GenerateVerification(userId)
	_, err = u.verificationRepo.Create(context, verification)
	if err != nil {
		return "", err
	}

	u.mailer.SendMail(string(*user.Email), verification.Code)

	return verification.Code, nil
	return "", nil
}
