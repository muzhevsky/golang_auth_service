package usecases

import (
	"authorization/internal"
	"authorization/internal/entities"
	errors2 "authorization/internal/errs"
	"authorization/internal/infrastructure/services/tokens"
	"context"
	"time"
)

type verifyTokenUseCase struct {
	sessionRepo  internal.ISessionRepository
	tokenManager tokens.ISessionManager
}

func NewVerifyTokenUseCase(sessionRepo internal.ISessionRepository, tokenManager tokens.ISessionManager) *verifyTokenUseCase {
	return &verifyTokenUseCase{sessionRepo: sessionRepo, tokenManager: tokenManager}
}

func (s *verifyTokenUseCase) VerifyAccessToken(context context.Context, token string) error {
	_, err := s.sessionRepo.FindByAccessToken(context, token)

	if err != nil {
		return errors2.NotAValidAccessToken
	}
	claimsMap, err := s.tokenManager.ParseToken(token)
	if claimsMap == nil {
		return errors2.NotAValidAccessToken
	}
	claims := entities.NewClaimsFromMap(claimsMap)
	if claims.ExpireAt.Before(time.Now()) {
		return errors2.AccessTokenExpired
	}

	return nil
}
