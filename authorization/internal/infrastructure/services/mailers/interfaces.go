package mailers

import "authorization/internal/entities/session"

type (
	IVerificationMailer interface {
		SendVerificationMail(email string, verificationCode string) error
	}

	INewSignInMailer interface {
		SendNewSignInMail(email string, device *session.Device) error
	}
)
