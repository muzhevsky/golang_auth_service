package usecases

import (
	"authorization/controllers/requests"
	"authorization/internal"
	"authorization/internal/entities/account"
	session2 "authorization/internal/entities/session"
	errors2 "authorization/internal/errs"
	"authorization/internal/infrastructure/services/mailers"
	tokens2 "authorization/internal/infrastructure/services/tokens"
	"context"
	"time"
)

type signInUseCase struct {
	userRepo        internal.IAccountRepository
	sessionRepo     internal.ISessionRepository
	deviceRepo      internal.IDeviceRepository
	newSignInMailer mailers.INewSignInMailer
	hashProvider    tokens2.IHashProvider
	sessionManager  tokens2.ISessionManager
}

func NewSignInUseCase(
	userRepo internal.IAccountRepository,
	sessionRepo internal.ISessionRepository,
	deviceRepo internal.IDeviceRepository,
	hashProvider tokens2.IHashProvider,
	sessionManager tokens2.ISessionManager,
	newSignInMailer mailers.INewSignInMailer) internal.ISignInUseCase {
	return &signInUseCase{
		userRepo:        userRepo,
		sessionRepo:     sessionRepo,
		deviceRepo:      deviceRepo,
		hashProvider:    hashProvider,
		sessionManager:  sessionManager,
		newSignInMailer: newSignInMailer,
	}
}

func (u *signInUseCase) SignIn(context context.Context, userRequest *requests.SignInRequest) (*requests.SignInResponse, error) {
	var accountRecord *account.Account
	login := account.Login(userRequest.Login)
	email := account.Email(userRequest.Login)

	accountRecord, err := u.userRepo.FindByLogin(context, login)
	if err != nil {
		return nil, err

	}
	if accountRecord == nil {
		accountRecord, err = u.userRepo.FindByEmail(context, email)
		if err != nil {
			return nil, err
		}
		if accountRecord == nil {
			return nil, errors2.AccountNotFound
		}
	}

	passwordMatched := u.hashProvider.CompareStringAndHash(userRequest.Password, string(accountRecord.Password))
	if !passwordMatched {
		return nil, errors2.WrongPassword
	}

	session, err := u.sessionManager.CreateSession(accountRecord)
	if err != nil {
		return nil, err
	}

	device := &session2.Device{
		AccountId:           accountRecord.Id,
		Name:                userRequest.DeviceName,
		SessionAccessToken:  session.AccessToken,
		SessionCreationTime: time.Now(),
	}

	go u.newSignInMailer.SendNewSignInMail(string(accountRecord.Email), device) // todo

	err = u.deviceRepo.Create(context, device)

	err = u.sessionRepo.Create(context, session)
	if err != nil {
		return nil, err
	}

	return &requests.SignInResponse{
		Id: accountRecord.Id,
		Session: requests.RefreshSessionResponse{
			AccessToken:  session.AccessToken,
			RefreshToken: session.RefreshToken,
			ExpiresAt:    session.AccessExpiresAt.Unix(),
		},
	}, nil
}
