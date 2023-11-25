package entities

type lengthValidator struct {
	minLength int
	maxLength int
}

func (v *lengthValidator) IsValid(login string) (bool, error) {
	if len(login) < v.minLength || len(login) > v.maxLength {
		return false, nil
	}
	return true, nil
}
