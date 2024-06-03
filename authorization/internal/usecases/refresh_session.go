package usecases

import (
	"authorization/internal"
	"authorization/internal/controllers/requests"
	"authorization/internal/entities"
	errors2 "authorization/internal/errs"
	"authorization/internal/infrastructure/services/tokens"
	"context"
)

type refreshSessionUseCase struct {
	userRepo     internal.IUserRepository
	sessionRepo  internal.ISessionRepository
	tokenManager tokens.ISessionManager
}

func NewRefreshSessionUseCase(userRepo internal.IUserRepository, sessionRepo internal.ISessionRepository, tokenManager tokens.ISessionManager) *refreshSessionUseCase {
	return &refreshSessionUseCase{userRepo: userRepo, sessionRepo: sessionRepo, tokenManager: tokenManager}
}

func (s *refreshSessionUseCase) RefreshSession(context context.Context, request *requests.RefreshSessionRequest) (*entities.Session, error) {
	accessToken := request.AccessToken
	refreshToken := request.RefreshToken
	err := s.verifyToken(request.AccessToken)
	if err != nil {
		return nil, err
	}

	storedSession, err := s.sessionRepo.FindByAccessToken(context, accessToken)
	if err != nil {
		return nil, err
	}

	if refreshToken != storedSession.RefreshToken {
		return nil, errors2.NotAValidRefreshToken
	}

	userId := storedSession.UserId
	user, err := s.userRepo.FindById(context, userId)
	if err != nil {
		return nil, err
	}

	newSession, err := s.tokenManager.CreateSession(user)
	if err != nil {
		return nil, err
	}

	result, err := s.sessionRepo.Update(context, storedSession, func(session *entities.Session) {
		session.ExpireAt = newSession.ExpireAt
		session.AccessToken = newSession.AccessToken
		session.RefreshToken = newSession.RefreshToken
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *refreshSessionUseCase) verifyToken(token string) error {
	claimsMap, err := s.tokenManager.ParseToken(token)
	if err != nil {
		return errors2.NotAValidAccessToken
	}
	claims := entities.NewClaimsFromMap(claimsMap)
	if claims == nil {
		return errors2.NotAValidAccessToken
	}

	return nil
}
