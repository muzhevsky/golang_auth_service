package usecase

import (
	"authorization/internal/entities"
	"context"
	"fmt"
	"time"
)

type userUseCase struct {
	userRepo         IUserRepo
	hashProvider     IPasswordHashProvider
	verificationRepo IVerificationRepo
}

func NewUser(
	userRepo IUserRepo,
	verificationRepo IVerificationRepo,
	hashProvider IPasswordHashProvider) *userUseCase {
	return &userUseCase{
		userRepo:         userRepo,
		verificationRepo: verificationRepo,
		hashProvider:     hashProvider,
	}
}
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
		return nil, fmt.Errorf("%w. login already exists", RecordAlreadyExists)
	}

	result, err = u.userRepo.CheckEmailExist(context, user.EMail)
	if err != nil {
		return nil, err
	}
	if result {
		return nil, fmt.Errorf("%w. email already exists", RecordAlreadyExists)
	}

	hashedPassword, err := u.hashProvider.GenerateHashPassword(user.Password)
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

func (u *userUseCase) SignIn(context context.Context, userRequest *entities.User) (result bool, err error) {
	userRecord, err := u.userRepo.FindOne(context, userRequest)
	if err != nil {
		return false, err
	}

	if !userRecord.IsVerified {
		return false, entities.UserIsNotVerified
	}

	passwordMatched := u.hashProvider.CompareStringAndHash(userRequest.Password, userRecord.Password)
	return passwordMatched, nil
}
