package usecases

import (
	"authorization/controllers/requests"
	"authorization/internal"
	accountpkg "authorization/internal/entities/account"
	session2 "authorization/internal/entities/session"
	"authorization/internal/errs"
	"authorization/internal/infrastructure/services/mailers"
	"authorization/internal/infrastructure/services/tokens"
	"context"
	"fmt"
	"time"
)

type createAccountUseCase struct {
	userRepo       internal.IAccountRepository
	hashProvider   tokens.IHashProvider
	sessionRepo    internal.ISessionRepository
	deviceRepo     internal.IDeviceRepository
	sessionManager tokens.ISessionManager
	mailer         mailers.IVerificationMailer
}

func NewCreateAccountUseCase(
	userRepo internal.IAccountRepository,
	sessionRepo internal.ISessionRepository,
	deviceRepo internal.IDeviceRepository,
	sessionManager tokens.ISessionManager,
	hashProvider tokens.IHashProvider,
	mailer mailers.IVerificationMailer) internal.ICreateAccountUseCase {
	return &createAccountUseCase{
		userRepo:       userRepo,
		sessionManager: sessionManager,
		sessionRepo:    sessionRepo,
		deviceRepo:     deviceRepo,
		hashProvider:   hashProvider,
		mailer:         mailer,
	}
}

// CreateAccount creates new account if it satisfies the necessary conditions and stores it using IAccountRepository
//
// Returns: requests.SignUpResponse
//
// Possible errors:
//   - errs.LoginValidationError, errs.EmailValidationError, errs.PasswordValidationError
//   - errs.RecordAlreadyExists if email or login are not unique
//   - errors of infrastructure from sources like IHashProvider or IAccountRepository implementations
func (u *createAccountUseCase) CreateAccount(context context.Context, request *requests.SignUpRequest) (*requests.SignUpResponse, error) {
	account := u.createAccount(request)
	err := account.Validate()
	if err != nil {
		return nil, err
	}

	exists, err := u.userRepo.CheckLoginExist(context, account.Login)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("%w. Login already exists", errs.RecordAlreadyExists)
	}

	exists, err = u.userRepo.CheckEmailExist(context, account.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("%w. Email already exists", errs.RecordAlreadyExists)
	}

	hashedPassword, err := u.hashProvider.GenerateHash(request.Password)
	if err != nil {
		return nil, err
	}

	account.ConfirmCreationBytes(hashedPassword)

	account.Id, err = u.userRepo.Create(context, account)
	if err != nil {
		return nil, err
	}

	session, err := u.sessionManager.CreateSession(account)
	if err != nil {
		return nil, err
	}

	device := u.createDevice(account, request, session)
	err = u.deviceRepo.Create(context, device)
	if err != nil {
		return nil, err
	}

	err = u.sessionRepo.Create(context, session)
	if err != nil {
		return nil, err
	}

	return u.createResponse(session, account), nil
}

func (u *createAccountUseCase) createAccount(request *requests.SignUpRequest) *accountpkg.Account {
	login := accountpkg.Login(request.Login)
	email := accountpkg.Email(request.Email)
	password := accountpkg.Password(request.Password)

	return &accountpkg.Account{
		Login:    login,
		Password: password,
		Email:    email,
	}
}

func (u *createAccountUseCase) createDevice(account *accountpkg.Account, request *requests.SignUpRequest, session *session2.Session) *session2.Device {
	return &session2.Device{
		AccountId:           account.Id,
		Name:                request.DeviceName,
		SessionAccessToken:  session.AccessToken,
		SessionCreationTime: time.Now(),
	}
}

func (u *createAccountUseCase) createResponse(session *session2.Session, account *accountpkg.Account) *requests.SignUpResponse {
	return &requests.SignUpResponse{
		Id: account.Id,
		Session: requests.RefreshSessionResponse{
			AccessToken:  session.AccessToken,
			RefreshToken: session.RefreshToken,
			ExpiresAt:    session.AccessExpiresAt.Unix(),
		},
	}
}
