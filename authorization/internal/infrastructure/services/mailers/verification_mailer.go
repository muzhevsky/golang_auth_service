package mailers

import (
	"fmt"
)

type verificationMailer struct {
	mailer *smtpMailer
}

func NewVerificationMailer(mailer *smtpMailer) *verificationMailer {
	return &verificationMailer{mailer: mailer}
}

func (vm *verificationMailer) SendMail(email string, verificationCode string) {
	vm.mailer.SendMail(email, "Код подтверждения", fmt.Sprintf(
		"Я могу поставить сюда другой текст, если будет нужно: %v", verificationCode))
}
