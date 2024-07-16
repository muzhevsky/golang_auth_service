package smtp

import (
	"log"
	"net/smtp"
)

type SMTP struct {
	auth     smtp.Auth
	host     string
	username string
	sender   string
	port     string
}

func New(username, sender, password, host, port string, opts ...Option) *SMTP {

	s := &SMTP{
		auth:     smtp.PlainAuth("", username, password, host),
		sender:   sender,
		port:     port,
		host:     host,
		username: username,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *SMTP) SendMail(to, subject, body string) error {
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"From: " + s.sender + "\r\n" +
		"\r\n" + body + "\r\n")
	addr := s.host + ":" + s.port
	err := smtp.SendMail(addr, s.auth, s.sender, []string{to}, msg)
	log.Printf(to + " " + subject + " " + body + " " + s.username + " " + s.host + " " + s.port)

	return err
}
