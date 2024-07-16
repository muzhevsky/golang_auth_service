package internal

import (
	"authorization/controllers/requests"
	"authorization/internal/entities/account"
	"authorization/internal/entities/session"
	"authorization/internal/entities/verification"
	"context"
)

type (
	ICreateAccountUseCase interface {
		CreateAccount(context context.Context, request *requests.SignUpRequest) (*requests.SignUpResponse, error)
	}

	IGetAccountDevicesUseCase interface {
		GetAccountDevices(context context.Context, accountId int) (*requests.AccountDevicesResponse, error)
	}

	ICloseSessionsByIdsUseCase interface {
		CloseSessionsByIds(context context.Context, accountId int, request *requests.CloseSessionsRequest) error
	}

	ISignInUseCase interface {
		SignIn(context context.Context, user *requests.SignInRequest) (*requests.SignInResponse, error)
	}

	IVerifyUserUseCase interface {
		Verify(context context.Context, userId int, code string) error
	}

	IRequestVerificationUseCase interface {
		RequestVerification(context context.Context, userId int) (string, error) // todo убрать код)
	}

	IRefreshSessionUseCase interface {
		RefreshSession(context context.Context, tokens *requests.RefreshSessionRequest) (*requests.RefreshSessionResponse, error)
	}

	ICheckVerificationUseCase interface {
		CheckVerification(context context.Context, accountId int) (bool, error)
	}

	IAccountRepository interface {
		Create(context context.Context, user *account.Account) (int, error)
		FindById(context context.Context, id int) (*account.Account, error)
		FindByLogin(context context.Context, login account.Login) (*account.Account, error)
		FindByEmail(context context.Context, email account.Email) (*account.Account, error)
		CheckLoginExist(context context.Context, login account.Login) (bool, error)
		CheckEmailExist(context context.Context, email account.Email) (bool, error)
		UpdateById(context context.Context, id int, user *account.Account) error
	}

	IVerificationRepository interface {
		Create(context context.Context, verification *verification.Verification) error
		FindByAccountId(context context.Context, userId int) ([]*verification.Verification, error)
		Clear(context context.Context, userId int) error
	}

	ISessionRepository interface {
		Create(context context.Context, user *session.Session) error
		FindByAccessToken(context context.Context, token string) (*session.Session, error)
		UpdateByAccessToken(context context.Context, token string, newSession *session.Session) (*session.Session, error)
		DeleteByAccessToken(context context.Context, token string) error
	}

	IDeviceRepository interface {
		Create(context context.Context, device *session.Device) error
		SelectByAccountId(context context.Context, accountId int) ([]*session.Device, error)
		SelectByAccessToken(context context.Context, token string) (*session.Device, error)
		SelectById(context context.Context, id int) (*session.Device, error)
		DeleteById(context context.Context, deviceId int) error
		UpdateByAccessToken(context context.Context, accessToken string, device *session.Device) error
	}
)
