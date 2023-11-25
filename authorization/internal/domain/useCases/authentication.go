package useCases

import (
	"authorization/internal/domain/entities"
	"authorization/internal/domain/errors"
	"log"
)

type authUseCase struct {
	repository        UserRepository
	encryptionService StringEncryptor
}

func NewAuthenticationService(
	repository UserRepository,
	encryptionService StringEncryptor) AuthenticationUseCase {
	return &authUseCase{repository, encryptionService}
}

func (useCase *authUseCase) SignUp(dto *entities.User) error {
	loginSelection, err := useCase.repository.SelectByLogin(dto.Login)
	if err != nil {
		log.Fatal("FATAL: couldn't execute repository query (authentication.go - SignUp())") // todo подумать над ошибками
		return err
	}
	if loginSelection != nil {
		return errors.LoginIsNotUnique
	}

	emailSelection, err := useCase.repository.SelectByEmail(dto.Email)
	if err != nil {
		return err
	}
	if emailSelection != nil {
		return errors.EmailIsNotUnique
	}

	hashedPassword, err := useCase.encryptionService.EncryptString(dto.Password)
	if err != nil {
		return err
	}

	user := entities.CreateUser(dto.Login, dto.Email, hashedPassword)

	err = useCase.repository.Insert(user)
	if err != nil {
		return err
	}

	return nil
}
