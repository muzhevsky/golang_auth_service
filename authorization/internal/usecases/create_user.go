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
	userRepo       internal.IUserRepository
	hashProvider   tokens.IHashProvider
	sessionRepo    internal.ISessionRepository
	sessionManager tokens.ISessionManager
	mailer         mailers.IVerificationMailer
}

func NewCreateUserUseCase(
	userRepo internal.IUserRepository,
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
func (u *userUseCase) CreateUser(context context.Context, request *requests.CreateUserRequest) (*requests.CreateUserResponse, error) {
	user := &entities.User{
		Login:    request.Login,
		Password: request.Password,
		EMail:    request.EMail,
		Nickname: request.Nickname,
	}

	err := validateFields(user)
	if err != nil {
		return nil, err
	}

	exists, err := u.userRepo.CheckLoginExist(context, user.Login)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("%w. Login already exists", errs.RecordAlreadyExists)
	}

	exists, err = u.userRepo.CheckEmailExist(context, user.EMail)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("%w. Email already exists", errs.RecordAlreadyExists)
	}

	hashedPassword, err := u.hashProvider.GenerateHash(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)
	user.CreationTime = time.Now()

	user.Id, err = u.userRepo.Create(context, user)
	if err != nil {
		return nil, err
	}

	session, err := u.sessionManager.CreateSession(user)
	if err != nil {
		return nil, err
	}

	_, err = u.sessionRepo.Create(context, session)
	if err != nil {
		return nil, err
	}

	return &requests.CreateUserResponse{
		Id: user.Id,
		Session: requests.RefreshSessionResponse{
			AccessToken:  session.AccessToken,
			RefreshToken: session.RefreshToken,
			ExpiresAt:    session.ExpiresAt.Unix(),
		},
	}, nil
}

func validateFields(user *entities.User) error {
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
