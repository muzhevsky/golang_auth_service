package mailers

import (
	"authorization/pkg/smtp"
	"fmt"
)

type smtpVerificationMailer struct {
	mailer *smtp.SMTP
}

func NewSMTPVerificationMailer(mailer *smtp.SMTP) IVerificationMailer {
	return &smtpVerificationMailer{mailer: mailer}
}

func (vm *smtpVerificationMailer) SendMail(email string, verificationCode string) {
	vm.mailer.SendMail(email, "Код подтверждения", bodyFromTemplate(verificationCode))
}

func bodyFromTemplate(verificationCode string) string {
	return fmt.Sprintf("here should be HTML-code to represent email properly, but who cares?\n"+
		"here's the code: %v", verificationCode)
}
