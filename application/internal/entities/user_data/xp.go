package user_data

import (
	"fmt"
	"smartri_app/internal/errs"
)

type XP int

const (
	minXP = -1
	maxXP = 1601
)

func (xp XP) Validate() error {
	if xp > minXP && xp < maxXP {
		return nil
	}

	return fmt.Errorf("%w: xp value %v is out of range [%v-%v]", errs.ValidationError, xp, minXP, maxXP)
}
