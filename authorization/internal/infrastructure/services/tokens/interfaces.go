package tokens

import "authorization/internal/entities"

type (
	IHashProvider interface {
		GenerateHash(stringToHash string) ([]byte, error)
		CompareStringAndHash(stringToCompare string, hashedString string) bool
	}

	ISessionManager interface {
		CreateSession(user *entities.User) (*entities.Session, error)
		ParseToken(token string) (map[string]interface{}, error)
	}

	IAccessTokenManager interface {
		CreateToken(claims map[string]interface{}) (string, error)
		ParseToken(token string) (map[string]interface{}, error)
	}

	IRefreshTokenGenerator interface {
		GenerateToken(userId int) (string, error)
	}
)
