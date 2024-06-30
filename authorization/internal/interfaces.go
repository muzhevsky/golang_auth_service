package internal

import (
	"authorization/internal/controllers/requests"
	"authorization/internal/entities"
	"context"
)

type (
	ICreateAccountUseCase interface {
		CreateAccount(context context.Context, request *requests.CreateAccountRequest) (*requests.CreateAccountResponse, error)
	}

	ISignInUseCase interface {
		SignIn(context context.Context, user *requests.SignInRequest) (*entities.Session, error)
	}

	IVerifyUserUseCase interface {
		Verify(context context.Context, userId int, code string) error
	}

	IRequestVerificationUseCase interface {
		RequestVerification(context context.Context, userId int) (string, error) // todo убрать код)
	}

	IRefreshSessionUseCase interface {
		RefreshSession(context context.Context, tokens *requests.RefreshSessionRequest) (*entities.Session, error)
	}

	ICheckVerificationUseCase interface {
		Check(context context.Context, accountId int) (bool, error)
	}

	IAccountRepository interface {
		Create(context context.Context, user *entities.Account) (int, error)
		FindById(context context.Context, id int) (*entities.Account, error)
		FindByLogin(context context.Context, login string) (*entities.Account, error)
		FindByEmail(context context.Context, email string) (*entities.Account, error)
		CheckLoginExist(context context.Context, login string) (bool, error)
		CheckEmailExist(context context.Context, email string) (bool, error)
		Update(context context.Context, user *entities.Account) error
	}

	IVerificationRepository interface {
		Create(context context.Context, verification *entities.Verification) (int, error)
		FindById(context context.Context, id int) (*entities.Verification, error)
		FindByAccountId(context context.Context, userId int) ([]*entities.Verification, error)
		Clear(context context.Context, userId int) error
	}

	ISessionRepository interface {
		Create(context context.Context, user *entities.Session) (int, error)
		FindByAccessToken(context context.Context, token string) (*entities.Session, error)
		Update(context context.Context, session *entities.Session, newSession *entities.Session) (*entities.Session, error)
	}
)
