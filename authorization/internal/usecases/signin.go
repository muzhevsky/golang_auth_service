package usecases

import (
	"authorization/internal"
	"authorization/internal/controllers/requests"
	"authorization/internal/entities"
	errors2 "authorization/internal/errs"
	tokens2 "authorization/internal/infrastructure/services/tokens"
	"context"
)

type signInUseCase struct {
	userRepo       internal.IUserRepository
	sessionRepo    internal.ISessionRepository
	hashProvider   tokens2.IHashProvider
	sessionManager tokens2.ISessionManager
}

func NewSignInUseCase(userRepo internal.IUserRepository, sessionRepo internal.ISessionRepository, hashProvider tokens2.IHashProvider, sessionManager tokens2.ISessionManager) *signInUseCase {
	return &signInUseCase{userRepo: userRepo, sessionRepo: sessionRepo, hashProvider: hashProvider, sessionManager: sessionManager}
}

func (u *signInUseCase) SignIn(context context.Context, userRequest *requests.SignInRequest) (*entities.Session, error) {
	var userRecord *entities.User
	userRecord, err := u.userRepo.FindByLogin(context, userRequest.Login)
	if err != nil {
		return nil, err

	}
	if userRecord == nil {
		userRecord, err = u.userRepo.FindByEmail(context, userRequest.Login)
		if err != nil {
			return nil, err
		}
		if userRecord == nil {
			return nil, errors2.UserNotFound
		}
	}

	passwordMatched := u.hashProvider.CompareStringAndHash(userRequest.Password, userRecord.Password)
	if !passwordMatched {
		return nil, errors2.WrongPassword
	}

	session, err := u.sessionManager.CreateSession(userRecord)
	if err != nil {
		return nil, err
	}

	_, err = u.sessionRepo.Create(context, session)
	if err != nil {
		return nil, err
	}

	return session, nil
}
