package usecase

import (
	"authorization/internal/entities"
	"context"
	"errors"
	"fmt"
	"time"
)

const (
	defaultAccessExpireDuration  = time.Duration(30) * time.Second
	defaultRefreshExpireDuration = time.Duration(2.592e15)
)

type SessionUseCase struct {
	accessExpireDuration    time.Duration
	refreshExpireDuration   time.Duration
	accessTokenManager      IAccessTokenManager
	refreshTokenGenerator   IRefreshTokenGenerator
	sessionRepo             ISessionRepo
	fingerprintHashProvider IHashProvider
}

func NewSessionUseCase(
	accessTokenManager IAccessTokenManager,
	refreshTokenManager IRefreshTokenGenerator,
	repo ISessionRepo,
	fingerprintHashProvider IHashProvider,
	options ...sessionUseCaseOption) *SessionUseCase {
	useCase := &SessionUseCase{
		defaultAccessExpireDuration,
		defaultRefreshExpireDuration,
		accessTokenManager,
		refreshTokenManager,
		repo,
		fingerprintHashProvider}

	for i := range options {
		options[i](useCase)
	}

	return useCase
}

func (s *SessionUseCase) VerifyAccessToken(context context.Context, token string) error {
	_, err := s.sessionRepo.FindByAccess(context, token)

	if err != nil {
		return entities.NotAValidAccessToken
	}
	claimsMap, err := s.accessTokenManager.ParseToken(token)
	if err != nil {
		return entities.NotAValidAccessToken
	}
	claims := entities.NewClaimsFromMap(claimsMap)
	if claims == nil {
		return entities.NotAValidAccessToken
	}
	if claims.ExpireAt.Before(time.Now()) {
		return entities.AccessTokenExpired
	}

	return nil
}

func (s *SessionUseCase) GetSession(context context.Context, token string) (*entities.Session, error) {
	session, err := s.sessionRepo.FindByAccess(context, token)
	if err != nil {
		return nil, entities.NotAValidAccessToken
	}
	return session, nil
}

func (s *SessionUseCase) CreateSession(context context.Context, user *entities.User) (*entities.Session, error) {
	result := &entities.Session{}
	accessToken, err := s.createAccess(user)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.createRefresh(user)
	if err != nil {
		return nil, err
	}
	fmt.Println(accessToken)
	result.UserId = user.Id
	result.DeviceDescription = /*"TODO"*/ ""
	result.AccessToken = accessToken
	result.RefreshToken = refreshToken
	result.ExpireAt = time.Now().Add(s.refreshExpireDuration)

	err = s.sessionRepo.Create(context, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SessionUseCase) UpdateSession(context context.Context, session *entities.Session) (*entities.Session, error) {
	accessToken := session.AccessToken
	refreshToken := session.RefreshToken
	err := s.VerifyAccessToken(context, session.AccessToken)
	if err != nil && !errors.Is(err, entities.AccessTokenExpired) {
		return nil, err
	}

	storedSession, err := s.GetSession(context, accessToken)
	if err != nil {
		return nil, err
	}

	if refreshToken != storedSession.RefreshToken {
		return nil, entities.NotAValidRefreshToken
	}
	user := &entities.User{}
	user.Id = storedSession.UserId

	err = s.sessionRepo.Delete(context, &entities.Session{user.Id, accessToken, refreshToken, "", time.Now()})
	if err != nil {
		return nil, err
	}
	return s.CreateSession(context, user)
}

func (s *SessionUseCase) createAccess(user *entities.User) (string, error) {
	claims := entities.NewClaims(user.Id, s.accessExpireDuration)
	token, err := s.accessTokenManager.GenerateToken(claims.MapFromClaims())
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *SessionUseCase) createRefresh(user *entities.User) (string, error) {
	token, err := s.refreshTokenGenerator.GenerateToken(user.Id)
	if err != nil {
		return "", err
	}
	return token, nil
}
