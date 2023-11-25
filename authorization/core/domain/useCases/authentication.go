package useCases

import (
	"authorization/core/domain/dtos"
	"authorization/core/domain/entities"
	"authorization/core/domain/errors"
	"authorization/core/domain/repositories/abstraction"
)

type AuthenticationService interface {
	SignUp(dto *dtos.SignUpDto) error
}

type authenticationService struct {
	repository        abstraction.UserRepository
	saltService       SaltGenerationService
	encryptionService PasswordEncryptionService
}

func NewAuthenticationService(
	repository abstraction.UserRepository,
	saltService SaltGenerationService,
	encryptionService PasswordEncryptionService) AuthenticationService {
	return &authenticationService{repository, saltService, encryptionService}
}

func (service *authenticationService) SignUp(dto *dtos.SignUpDto) error {
	loginSelection, err := service.repository.SelectByLogin(dto.Login)
	if err != nil {
		return err
	}
	if loginSelection != nil {
		return &errors.NotUniqueLogin{}
	}

	emailSelection, err := service.repository.SelectByEmail(dto.Email)
	if err != nil {
		return err
	}
	if emailSelection != nil {
		return &errors.NotUniqueEmail{}
	}

	user := entities.User{}

	hashedPassword := dto.Password
	hashedPassword += service.saltService.GenerateSalt()
	hashedPassword, err = service.encryptionService.Encrypt(hashedPassword)
	if err != nil {
		return err
	}

	user.Login = dto.Login
	user.Email = dto.Email
	user.Password = hashedPassword

	err = service.repository.Insert(&user)
	if err != nil {
		return err
	}

	return nil
}
