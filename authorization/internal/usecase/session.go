package usecase

import (
	"authorization/internal/entities"
	"context"
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
	userRepo              IUserRepo
}

func NewSessionUseCase(
	accessTokenManager ITokenManager,
	refreshTokenManager ITokenManager,
	repo ISessionRepo,
	userRepo IUserRepo,
	options ...sessionUseCaseOption) *SessionUseCase {
	useCase := &SessionUseCase{defaultAccessExpireDuration, defaultRefreshExpireDuration,
		accessTokenManager, refreshTokenManager, repo, userRepo}

	for i := range options {
		options[i](useCase)
	}

	return useCase
}

func (s *SessionUseCase) VerifyAccessToken(context context.Context, token string) (bool, error) {
	claimsMap, err := s.accessTokenManager.ParseToken(token)
	if err != nil {
		return false, entities.NotAValidToken
	}
	claims := entities.NewClaimsFromMap(claimsMap)
	if claims.ExpireAt.Before(time.Now()) {
		return false, entities.TokenExpired
	}

	return true, nil
}

func (s *SessionUseCase) CreateTokens(context context.Context, user *entities.User) (*entities.Session, error) {
	result := &entities.Session{}
	accessToken, err := s.createAccess(user)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.createRefresh()
	if err != nil {
		return nil, err
	}

	result.AccessToken = accessToken
	result.RefreshToken = refreshToken
	result.ExpireAt = time.Now().Add(s.refreshExpireDuration)

	err = s.sessionRepo.Create(context, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *SessionUseCase) UpdateAccessToken(context context.Context, accessToken, refreshToken string) (*entities.Session, error) {
	_, err := s.refreshTokenManager.ParseToken(refreshToken)
	if err != nil {
		return nil, err
	}

	claims, err := s.accessTokenManager.ParseToken(accessToken)
	user, err := s.userRepo.FindById(context, claims["userId"].(int))
	if err != nil {
		return nil, err
	}

	err = s.sessionRepo.Delete(context, &entities.Session{accessToken, refreshToken, time.Now()})
	if err != nil {
		return nil, err
	}
	return s.CreateTokens(context, user)
}

func (s *SessionUseCase) createAccess(user *entities.User) (string, error) {
	claims := entities.NewClaims(user.Id, s.accessExpireDuration)
	token, err := s.accessTokenManager.GenerateToken(claims.MapFromClaims())
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *SessionUseCase) createRefresh() (string, error) {
	token, err := s.refreshTokenManager.GenerateToken(make(map[string]interface{}))
	if err != nil {
		return "", err
	}
	return token, nil
}
