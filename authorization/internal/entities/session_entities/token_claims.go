package session_entities

import (
	"authorization/internal/errs"
	"time"
)

const (
	IssuerName         = "iss"
	AccountIdClaimName = "accountId"
	ExpiresAt          = "expiresAt"
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

func NewClaimsFromMap(claimsMap map[string]any) (*TokenClaims, error) {
	userId, exists := claimsMap[AccountIdClaimName].(float64)
	if !exists {
		return nil, errs.NotAValidAccessToken
	}
	expiresAt, exists := claimsMap[ExpiresAt].(float64)
	if !exists {
		return nil, errs.NotAValidAccessToken
	}
	issuer, exists := claimsMap[IssuerName].(string)
	if !exists {
		return nil, errs.NotAValidAccessToken
	}

	return &TokenClaims{
		UserId:    int(userId),
		ExpiresAt: time.Unix(int64(expiresAt), 0),
		Issuer:    issuer}, nil
}

func (claims *TokenClaims) MapFromClaims() map[string]any {
	result := make(map[string]interface{})
	result[AccountIdClaimName] = claims.UserId
	result[ExpiresAt] = claims.ExpiresAt.Unix()
	result[IssuerName] = claims.Issuer

	return result
}
