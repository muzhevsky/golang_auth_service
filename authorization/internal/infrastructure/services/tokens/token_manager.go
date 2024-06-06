package tokens

import (
	"authorization/internal/entities"
	"time"
)

type tokenManager struct {
	config  TokenConfiguration
	access  IAccessTokenManager
	refresh IRefreshTokenGenerator
}

func NewTokenManager(config TokenConfiguration, access IAccessTokenManager, refresh IRefreshTokenGenerator) *tokenManager {
	return &tokenManager{config: config, access: access, refresh: refresh}
}

func (t *tokenManager) CreateSession(user *entities.User) (*entities.Session, error) {
	claims := make(map[string]interface{})
	expiresAt := time.Now().Add(t.config.AccessTokenDuration)
	claims["iss"] = t.config.Issuer
	claims["userId"] = user.Id
	claims["expiresAt"] = expiresAt.Unix()

	access, err := t.access.CreateToken(claims)
	if err != nil {
		return nil, err
	}

	refresh, err := t.refresh.GenerateToken(user.Id)
	if err != nil {
		return nil, err
	}

	return &entities.Session{
		UserId:         user.Id,
		DeviceIdentity: "TODO",
		AccessToken:    access,
		RefreshToken:   refresh,
		ExpiresAt:      expiresAt,
	}, nil
}

func (t *tokenManager) ParseToken(token string) (map[string]interface{}, error) {
	return t.access.ParseToken(token)
}
