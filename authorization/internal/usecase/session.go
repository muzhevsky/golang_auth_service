package usecase

import (
	"authorization/internal/entities"
	"time"
)

const (
	defaultAccessExpireDuration  = time.Duration(1.8e12)
	defaultRefreshExpireDuration = time.Duration(2.592e15)
)

type SessionUseCase struct {
	accessExpireDuration  time.Duration
	refreshExpireDuration time.Duration
	accessTokenManager    ITokenManager
	refreshTokenManager   ITokenManager
	sessionRepo           ISessionRepo
}

func NewSessionUseCase(accessTokenManager ITokenManager, refreshTokenManager ITokenManager, repo ISessionRepo, options ...sessionUseCaseOption) *SessionUseCase {
	useCase := &SessionUseCase{defaultAccessExpireDuration, defaultRefreshExpireDuration,
		accessTokenManager, refreshTokenManager, repo}

	for i := range options {
		options[i](useCase)
	}

	return useCase
}

func (s SessionUseCase) VerifyAccessToken(token string) (bool, error) {
	claimsMap, err := s.accessTokenManager.ParseToken(token)
	if err != nil {
		return false, err
	}
	claims := entities.NewClaimsFromMap(claimsMap)
	if claims.ExpireAt.Before(time.Now()) {
		return false, entities.TokenExpired
	}

	return true, nil
}

func (s SessionUseCase) CreateTokens(user *entities.User) (*entities.Session, error) {
	result := &entities.Session{}
	claims := entities.NewClaims(user.Id, time.Now().Add(s.accessExpireDuration))
	token, err := s.accessTokenManager.GenerateToken(claims.MapFromClaims())
	if err != nil {
		return nil, err
	}
	result.AccessToken = token

	token, err = s.refreshTokenManager.GenerateToken(make(map[string]interface{}))
	if err != nil {
		return nil, err
	}
	result.RefreshToken = token
	result.ExpireAt = time.Now().Add(s.refreshExpireDuration)

	return result, nil
}

func (s SessionUseCase) UpdateAccessToken(accessToken, refreshToken string) (*entities.Session, error) {
	//TODO implement me
	panic("implement me")
}
