package user_data_entities

import (
	"fmt"
	"smartri_app/internal/errs"
)

const (
	female = byte('f')
	male   = byte('m')
)

type Gender string

func (gender Gender) Validate() error {
	if (gender == "" || len(gender) > 1) || gender[0] != female && gender[0] != male {
		return fmt.Errorf("%w gender can only be 'f' or 'm'", errs.ValidationError)
	}

	return nil
}
