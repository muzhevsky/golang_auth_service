package tokens

import (
	"authorization/internal/entities/account"
	"authorization/internal/entities/session"
)

type (
	IHashProvider interface {
		GenerateHash(stringToHash string) ([]byte, error)
		CompareStringAndHash(stringToCompare string, hashedString string) bool
	}

	ISessionManager interface {
		CreateSession(user *account.Account) (*session.Session, error)
		ParseToken(token string) (*session.TokenClaims, error)
	}

	IAccessTokenManager interface {
		CreateToken(claims map[string]interface{}) (string, error)
		ParseToken(token string) (map[string]interface{}, error)
	}

	IRefreshTokenGenerator interface {
		GenerateToken(userId int) (string, error)
	}
)
