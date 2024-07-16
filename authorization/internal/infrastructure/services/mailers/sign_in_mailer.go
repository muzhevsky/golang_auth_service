package mailers

import (
	"authorization/internal/entities/session"
	"authorization/pkg/smtp"
	"fmt"
)

type signInMailer struct {
	mailer *smtp.SMTP
}

func NewSignInMailer(mailer *smtp.SMTP) INewSignInMailer {
	return &signInMailer{mailer: mailer}
}

func (vm *signInMailer) SendNewSignInMail(email string, device *session.Device) error {
	return vm.mailer.SendMail(email, "Новый вход в аккаунт", vm.bodyFromTemplate(device))
}

func (vm *signInMailer) bodyFromTemplate(device *session.Device) string {
	return fmt.Sprintf("here should be HTML-code to represent email properly, but who cares?\n"+
		"Кто-то вошел в аккаунт %v с нового устройства: %v", device.SessionCreationTime, device.Name)
}
