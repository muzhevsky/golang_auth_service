package mailers

type (
	IVerificationMailer interface {
		SendMail(email string, verificationCode string)
	}
)
