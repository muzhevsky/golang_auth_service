package usecases

import (
	"authorization/controllers/requests"
	"authorization/internal"
	"authorization/internal/entities/entities_account"
	session2 "authorization/internal/entities/session_entities"
	errors2 "authorization/internal/errs"
	"authorization/internal/infrastructure/services/mailers"
	tokens2 "authorization/internal/infrastructure/services/tokens"
	"context"
)

type signInUseCase struct {
	accountRepo     internal.IAccountRepository
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
		accountRepo:     userRepo,
		sessionRepo:     sessionRepo,
		deviceRepo:      deviceRepo,
		hashProvider:    hashProvider,
		sessionManager:  sessionManager,
		newSignInMailer: newSignInMailer,
	}
}

func (u *signInUseCase) SignIn(context context.Context, request *requests.SignInRequest) (*requests.SignInResponse, error) {
	var accountRecord *entities_account.Account

	accountRecord, err := u.findAccount(context, request.Login)
	if err != nil {
		return nil, err
	}

	passwordMatched := u.hashProvider.CompareStringAndHash(request.Password, string(accountRecord.Password))
	if !passwordMatched {
		return nil, errors2.WrongPassword
	}

	session, err := u.sessionManager.CreateSession(accountRecord)
	if err != nil {
		return nil, err
	}

	device := session2.NewDevice(accountRecord.Id, request.DeviceName, session.AccessToken)

	go u.newSignInMailer.SendNewSignInMail(string(accountRecord.Email), device) // todo

	err = u.deviceRepo.Create(context, device)
	if err != nil {
		return nil, err
	}

	err = u.sessionRepo.Create(context, session)
	if err != nil {
		return nil, err
	}

	refreshSessionResponse := requests.NewRefreshSessionResponse(session.AccessToken, session.RefreshToken, session.AccessExpiresAt.Unix())

	return requests.NewSignInResponse(accountRecord.Id, refreshSessionResponse), nil
}

func (u *signInUseCase) findAccount(context context.Context, credential string) (*entities_account.Account, error) {
	login := entities_account.Login(credential)
	acc, err := u.accountRepo.FindByLogin(context, login)
	if err != nil {
		return nil, err

	}
	if acc == nil {
		email := entities_account.Email(credential)
		acc, err = u.accountRepo.FindByEmail(context, email)
		if err != nil {
			return nil, err
		}
		if acc == nil {
			return nil, errors2.AccountNotFound
		}
	}

	return acc, nil
}
