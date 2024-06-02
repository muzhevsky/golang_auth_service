package tokens

import "time"

type TokenConfiguration struct {
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
	Issuer               string
	// todo fields
}
