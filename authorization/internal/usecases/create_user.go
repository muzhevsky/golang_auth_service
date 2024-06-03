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
	userRepo         internal.IUserRepository
	hashProvider     tokens.IHashProvider
	verificationRepo internal.IVerificationRepository
	mailer           mailers.IVerificationMailer
}

func NewCreateUserUseCase(
	userRepo internal.IUserRepository,
	verificationRepo internal.IVerificationRepository,
	hashProvider tokens.IHashProvider,
	mailer mailers.IVerificationMailer) *userUseCase {
	return &userUseCase{
		userRepo:         userRepo,
		verificationRepo: verificationRepo,
		hashProvider:     hashProvider,
		mailer:           mailer,
	}
}

// CreateUser - creates new record in database with user's repositories
// returns objects including user's repositories
// possible errors:
//   - validation errors
//   - non-unique login
//   - non-unique email
//   - errors of password hash and user repository
func (u *userUseCase) CreateUser(context context.Context, request *requests.CreateUserRequest) (*entities.User, error) {
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

	verification := entities.GenerateVerification(user.Id)
	_, err = u.verificationRepo.Create(context, verification)
	if err != nil {
		return nil, err
	}

	u.mailer.SendMail(user.EMail, verification.Code)

	return user, nil
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
