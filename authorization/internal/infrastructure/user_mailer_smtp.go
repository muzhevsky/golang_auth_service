package infrastructure

import (
	"authorization/internal/entities"
	"authorization/pkg/smtp"
)

type smtpMailer struct {
	smtp *smtp.SMTP
}

func NewSmtpMailer(smtp *smtp.SMTP) *smtpMailer {
	return &smtpMailer{smtp: smtp}
}

func (mailer *smtpMailer) SendMail(receiver string, subject string, body entities.EmailBody) {
	mailer.smtp.SendMail(
		receiver,
		subject,
		body.Body)
}
