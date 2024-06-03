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
	UserId   int
	ExpireAt time.Time
	Issuer   string
}

func NewClaims(userId int, duration time.Duration, issuer string) *TokenClaims {
	expireAt := time.Now().Add(duration)
	return &TokenClaims{UserId: userId, ExpireAt: expireAt, Issuer: issuer}
}

func NewClaimsFromMap(claimsMap map[string]interface{}) *TokenClaims {
	userId, exists := claimsMap[UserIdClaimName].(float64)
	if !exists {
		return nil
	}
	expireAt, exists := claimsMap[ExpiresAt].(float64)
	if !exists {
		return nil
	}
	issuer, exists := claimsMap[IssuerName].(string)
	if !exists {
		return nil
	}

	return &TokenClaims{int(userId), time.Unix(int64(expireAt), 0), issuer}
}

func (claims *TokenClaims) MapFromClaims() map[string]interface{} {
	result := make(map[string]interface{})
	result[UserIdClaimName] = claims.UserId
	result[ExpiresAt] = claims.ExpireAt.Unix()

	return result
}
