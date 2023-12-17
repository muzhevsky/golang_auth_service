package smtp

type Option func(*SMTP)

func Port(port string) Option {
	return func(s *SMTP) {
		s.port = port
	}
}
func Sender(sender string) Option {
	return func(s *SMTP) {
		s.sender = sender
	}
}
