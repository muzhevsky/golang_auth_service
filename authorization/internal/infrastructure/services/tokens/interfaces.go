package tokens

import (
	"authorization/internal/entities/entities_account"
	"authorization/internal/entities/session_entities"
)

type (
	IHashProvider interface {
		GenerateHash(stringToHash string) ([]byte, error)
		CompareStringAndHash(stringToCompare string, hashedString string) bool
	}

	ISessionManager interface {
		CreateSession(user *entities_account.Account) (*session_entities.Session, error)
		ParseToken(token string) (*session_entities.TokenClaims, error)
	}

	IAccessTokenManager interface {
		CreateToken(claims map[string]interface{}) (string, error)
		ParseToken(token string) (map[string]interface{}, error)
	}

	IRefreshTokenGenerator interface {
		GenerateToken(userId int) (string, error)
	}
)
