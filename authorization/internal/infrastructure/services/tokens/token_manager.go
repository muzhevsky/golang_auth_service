package tokens

import (
	"authorization/config"
	"authorization/internal/entities/entities_account"
	"authorization/internal/entities/session_entities"
	"time"
)

type tokenManager struct {
	config  config.TokenConfiguration
	access  IAccessTokenManager
	refresh IRefreshTokenGenerator
}

func NewTokenManager(
	config config.TokenConfiguration,
	access IAccessTokenManager,
	refresh IRefreshTokenGenerator) ISessionManager {
	return &tokenManager{config: config, access: access, refresh: refresh}
}

func (t *tokenManager) CreateSession(account *entities_account.Account) (*session_entities.Session, error) {
	refreshExpiresAt := time.Now().Add(time.Duration(int64(time.Second) * t.config.RefreshTokenDuration))

	claims := session_entities.NewClaims(account.Id, time.Duration(int64(time.Second)*t.config.AccessTokenDuration), t.config.Issuer)

	access, err := t.access.CreateToken(claims.MapFromClaims())
	if err != nil {
		return nil, err
	}

	refresh, err := t.refresh.GenerateToken(account.Id)
	if err != nil {
		return nil, err
	}

	return &session_entities.Session{
		AccountId:       account.Id,
		AccessToken:     access,
		AccessExpiresAt: claims.ExpiresAt,
		RefreshToken:    refresh,
		ExpiresAt:       refreshExpiresAt,
	}, nil
}

func (t *tokenManager) ParseToken(token string) (*session_entities.TokenClaims, error) {
	dict, err := t.access.ParseToken(token)
	if err != nil {
		return nil, err
	}
	claims, err := session_entities.NewClaimsFromMap(dict)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
