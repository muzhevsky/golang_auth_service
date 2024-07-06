package usecases

import (
	"authorization/controllers/requests"
	"authorization/internal"
	accountpkg "authorization/internal/entities/account"
	"authorization/internal/errs"
	"authorization/internal/infrastructure/services/mailers"
	"authorization/internal/infrastructure/services/tokens"
	"context"
	"fmt"
)

type userUseCase struct {
	userRepo       internal.IAccountRepository
	hashProvider   tokens.IHashProvider
	sessionRepo    internal.ISessionRepository
	sessionManager tokens.ISessionManager
	mailer         mailers.IVerificationMailer
}

func NewCreateUserUseCase(
	userRepo internal.IAccountRepository,
	sessionRepo internal.ISessionRepository,
	sessionManager tokens.ISessionManager,
	hashProvider tokens.IHashProvider,
	mailer mailers.IVerificationMailer) *userUseCase {
	return &userUseCase{
		userRepo:       userRepo,
		sessionManager: sessionManager,
		sessionRepo:    sessionRepo,
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
func (u *userUseCase) CreateAccount(context context.Context, request *requests.SignUpRequest) (*requests.SignUpResponse, error) {
	login := accountpkg.Login(request.Login)
	email := accountpkg.Email(request.Email)
	password := accountpkg.Password(request.Password)
	nickname := accountpkg.Nickname(login)

	account := &accountpkg.Account{
		Login:    login,
		Password: password,
		Email:    email,
		Nickname: nickname,
	}

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

	_, err = u.sessionRepo.Create(context, session)
	if err != nil {
		return nil, err
	}

	return &requests.SignUpResponse{
		Id: account.Id,
		Session: requests.RefreshSessionResponse{
			AccessToken:  session.AccessToken,
			RefreshToken: session.RefreshToken,
			ExpiresAt:    session.AccessExpiresAt.Unix(),
		},
	}, nil
}
