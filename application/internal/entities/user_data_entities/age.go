package user_data_entities

import (
	"fmt"
	"smartri_app/internal/errs"
)

const (
	minAge = 3
	maxAge = 100
)

type Age int

func (age Age) Validate() error {
	if minAge < age && age < maxAge {
		return nil
	}

	return fmt.Errorf("%w age can be in %d and %d bounds", errs.ValidationError, minAge, maxAge)
}
