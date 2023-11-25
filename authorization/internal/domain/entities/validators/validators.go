package entities

type LengthValidator interface {
	IsValid(string) (bool, error)
}
