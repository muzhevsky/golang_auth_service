package errors

type NotUniqueLogin struct{}

func (m *NotUniqueLogin) Error() string {
	return "user with specified login already exists"
}

type NotUniqueEmail struct{}

func (m *NotUniqueEmail) Error() string {
	return "user with specified email already exists"
}
