package entities

import (
	"time"
)

const (
	UserIdClaimName = "userId"
	ExpireAt        = "expireAt"
)

type TokenClaims struct {
	UserId   int
	ExpireAt time.Time
}

func NewClaims(userId int, duration time.Duration) *TokenClaims {
	expireAt := time.Now().Add(duration)
	return &TokenClaims{UserId: userId, ExpireAt: expireAt}
}

func NewClaimsFromMap(claimsMap map[string]interface{}) *TokenClaims {
	userId, exists := claimsMap[UserIdClaimName].(float64)
	if !exists {
		return nil
	}
	expireAt, exists := claimsMap[ExpireAt].(float64)
	if !exists {
		return nil
	}

	return &TokenClaims{int(userId), time.Unix(int64(expireAt), 0)}
}

func (claims *TokenClaims) MapFromClaims() map[string]interface{} {
	result := make(map[string]interface{})
	result[UserIdClaimName] = claims.UserId
	result[ExpireAt] = claims.ExpireAt.Unix()

	return result
}
