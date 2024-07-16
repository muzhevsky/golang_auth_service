package usecases

import (
	"authorization/controllers/requests"
	"authorization/internal"
	errors2 "authorization/internal/errs"
	"authorization/internal/infrastructure/services/tokens"
	"context"
)

type refreshSessionUseCase struct {
	accountRepo  internal.IAccountRepository
	sessionRepo  internal.ISessionRepository
	tokenManager tokens.ISessionManager
}

func NewRefreshSessionUseCase(accountRepo internal.IAccountRepository, sessionRepo internal.ISessionRepository, tokenManager tokens.ISessionManager) internal.IRefreshSessionUseCase {
	return &refreshSessionUseCase{accountRepo: accountRepo, sessionRepo: sessionRepo, tokenManager: tokenManager}
}

func (s *refreshSessionUseCase) RefreshSession(context context.Context, request *requests.RefreshSessionRequest) (*requests.RefreshSessionResponse, error) {
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

	userId := storedSession.AccountId
	user, err := s.accountRepo.FindById(context, userId)
	if err != nil {
		return nil, err
	}

	newSession, err := s.tokenManager.CreateSession(user)
	if err != nil {
		return nil, err
	}

	_, err = s.sessionRepo.UpdateByAccessToken(context, storedSession.AccessToken, newSession)
	if err != nil {
		return nil, err
	}

	return &requests.RefreshSessionResponse{
		AccessToken:  newSession.AccessToken,
		RefreshToken: newSession.RefreshToken,
		ExpiresAt:    newSession.ExpiresAt.Unix(),
	}, nil
}

func (s *refreshSessionUseCase) verifyToken(token string) error {
	_, err := s.tokenManager.ParseToken(token)
	if err != nil {
		return errors2.NotAValidAccessToken
	}

	return nil
}
