package entities

import (
	"time"
)

const (
	IssuerName      = "iss"
	UserIdClaimName = "userId"
	ExpiresAt       = "expiresAt"
)

type TokenClaims struct {
	UserId    int
	ExpiresAt time.Time
	Issuer    string
}

func NewClaims(userId int, duration time.Duration, issuer string) *TokenClaims {
	expiresAt := time.Now().Add(duration)
	return &TokenClaims{UserId: userId, ExpiresAt: expiresAt, Issuer: issuer}
}

func NewClaimsFromMap(claimsMap map[string]interface{}) *TokenClaims {
	userId, exists := claimsMap[UserIdClaimName].(float64)
	if !exists {
		return nil
	}
	expiresAt, exists := claimsMap[ExpiresAt].(float64)
	if !exists {
		return nil
	}
	issuer, exists := claimsMap[IssuerName].(string)
	if !exists {
		return nil
	}

	return &TokenClaims{int(userId), time.Unix(int64(expiresAt), 0), issuer}
}

func (claims *TokenClaims) MapFromClaims() map[string]interface{} {
	result := make(map[string]interface{})
	result[UserIdClaimName] = claims.UserId
	result[ExpiresAt] = claims.ExpiresAt.Unix()

	return result
}
