package internal

import (
	"authorization/internal/controllers/requests"
	"authorization/internal/entities"
	"context"
)

type (
	ICreateUserUseCase interface {
		CreateUser(context context.Context, user *requests.CreateUserRequest) (*entities.User, error)
	}

	ISignInUseCase interface {
		SignIn(context context.Context, user *requests.SignInRequest) (*entities.Session, error)
	}

	IVerifyUserUseCase interface {
		Verify(context context.Context, verification *requests.VerificationRequest) error
	}

	IRefreshSessionUseCase interface {
		RefreshSession(context context.Context, tokens *requests.RefreshSessionRequest) (*entities.Session, error)
	}

	IUserRepository interface {
		Create(context context.Context, user *entities.User) (int, error)
		FindById(context context.Context, id int) (*entities.User, error)
		FindByLogin(context context.Context, login string) (*entities.User, error)
		FindByEmail(context context.Context, email string) (*entities.User, error)
		CheckLoginExist(context context.Context, login string) (bool, error)
		CheckEmailExist(context context.Context, email string) (bool, error)
		Verify(context context.Context, id int) error
	}

	IVerificationRepository interface {
		Create(context context.Context, verification *entities.Verification) (int, error)
		FindById(context context.Context, id int) (*entities.Verification, error)
		FindByUserId(context context.Context, userId int) ([]*entities.Verification, error)
		Clear(context context.Context, userId int) error
	}

	ISessionRepository interface {
		Create(context context.Context, user *entities.Session) (int, error)
		FindByAccessToken(context context.Context, token string) (*entities.Session, error)
		Update(context context.Context, session *entities.Session, updateFunc func(session *entities.Session)) (*entities.Session, error)
	}
)
