package entities

import (
	"time"
)

const (
	UserId   = "userId"
	ExpireAt = "expireAt"
)

type TokenClaims struct {
	UserId   int
	ExpireAt time.Time
}

func NewClaims(userId int, expireAt time.Time) *TokenClaims {
	return &TokenClaims{UserId: userId, ExpireAt: expireAt}
}

func NewClaimsFromMap(claimsMap map[string]interface{}) *TokenClaims {
	userId := claimsMap[UserId].(int)
	expireAt := claimsMap[ExpireAt].(int64)
	return &TokenClaims{userId, time.Unix(expireAt, 0)}
}

func (claims *TokenClaims) MapFromClaims() map[string]interface{} {
	result := make(map[string]interface{})
	result[UserId] = claims.UserId
	result[ExpireAt] = claims.ExpireAt

	return result
}
