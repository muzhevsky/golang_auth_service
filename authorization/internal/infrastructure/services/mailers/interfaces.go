package mailers

import "authorization/internal/entities/session_entities"

type (
	IVerificationMailer interface {
		SendVerificationMail(email string, verificationCode string) error
	}

	INewSignInMailer interface {
		SendNewSignInMail(email string, device *session_entities.Device) error
	}
)
