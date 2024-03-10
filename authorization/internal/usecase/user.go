package usecase

import (
	"authorization/internal/entities"
	"context"
	"errors"
	"fmt"
	"time"
)

type userUseCase struct {
	userRepo         IUserRepo
	hashProvider     IHashProvider
	verificationRepo IVerificationRepo
}

func NewUser(
	userRepo IUserRepo,
	verificationRepo IVerificationRepo,
	hashProvider IHashProvider) *userUseCase {
	return &userUseCase{
		userRepo:         userRepo,
		verificationRepo: verificationRepo,
		hashProvider:     hashProvider,
	}
}

// CreateUser - creates new record in database with user's data
// returns objects including user's data
// possible errors:
//   - validation errors
//   - non-unique login
//   - non-unique email
//   - errors of password hash and user repository
func (u *userUseCase) CreateUser(context context.Context, user *entities.User) (*entities.User, error) {
	err := validateFields(user)
	if err != nil {
		return nil, err
	}

	result, err := u.userRepo.CheckLoginExist(context, user.Login)
	if err != nil {
		return nil, err
	}
	if result {
		return nil, fmt.Errorf("%w. Login already exists", RecordAlreadyExists)
	}

	result, err = u.userRepo.CheckEmailExist(context, user.EMail)
	if err != nil {
		return nil, err
	}
	if result {
		return nil, fmt.Errorf("%w. Email already exists", RecordAlreadyExists)
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

func (u *userUseCase) SignIn(context context.Context, userRequest *entities.User) (*entities.User, error) {
	var userRecord *entities.User
	userRecord, err := u.userRepo.FindByLogin(context, userRequest.Login)
	if err != nil {
		if errors.Is(err, entities.UserNotFound) {
			userRecord, err = u.userRepo.FindByEmail(context, userRequest.Login)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	if !userRecord.IsVerified {
		return nil, entities.UserIsNotVerified
	}

	passwordMatched := u.hashProvider.CompareStringAndHash(userRequest.Password, userRecord.Password)
	if !passwordMatched {
		return nil, entities.WrongPassword
	}
	return userRecord, nil
}
