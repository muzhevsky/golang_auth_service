package usecases

import (
	"authorization/internal"
	"authorization/internal/controllers/requests"
	"authorization/internal/entities"
	"authorization/internal/errs"
	"authorization/internal/infrastructure/services/mailers"
	"authorization/internal/infrastructure/services/tokens"
	"context"
	"fmt"
	"time"
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

// CreateUser - creates new record in database with user's repositories
// returns objects including user's repositories
// possible errors:
//   - validation errors
//   - non-unique login
//   - non-unique email
//   - errors of password hash and user repository
func (u *userUseCase) CreateAccount(context context.Context, request *requests.CreateAccountRequest) (*requests.CreateAccountResponse, error) {
	account := &entities.Account{
		Login:    request.Login,
		Password: request.Password,
		EMail:    request.EMail,
		Nickname: request.Nickname,
	}

	err := validateFields(account)
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

	exists, err = u.userRepo.CheckEmailExist(context, account.EMail)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("%w. Email already exists", errs.RecordAlreadyExists)
	}

	hashedPassword, err := u.hashProvider.GenerateHash(account.Password)
	if err != nil {
		return nil, err
	}

	account.Password = string(hashedPassword)
	account.CreationTime = time.Now()

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

	return &requests.CreateAccountResponse{
		Id: account.Id,
		Session: requests.RefreshSessionResponse{
			AccessToken:  session.AccessToken,
			RefreshToken: session.RefreshToken,
			ExpiresAt:    session.ExpiresAt.Unix(),
		},
	}, nil
}

func validateFields(user *entities.Account) error {
	validator := entities.UserValidator{}
	err := validator.ValidateLogin(user.Login)
	if err != nil {
		return err
	}
	err = validator.ValidatePassword(user.Password)
	if err != nil {
		return err
	}
	err = validator.ValidateEmail(user.EMail)
	if err != nil {
		return err
	}
	err = validator.ValidateNickname(user.Nickname)
	if err != nil {
		return err
	}
	return nil
}
