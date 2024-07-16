package datasources

import (
	"authorization/internal/entities/entities_account"
	"authorization/internal/entities/session_entities"
	"authorization/internal/entities/verification_entities"
	"context"
)

type (
	IInsertAccountCommand interface {
		Execute(context context.Context, user *entities_account.Account) (int, error)
	}
	ISelectAccountByIdCommand interface {
		Execute(context context.Context, id int) (*entities_account.Account, error)
	}
	ISelectAccountByLoginCommand interface {
		Execute(context context.Context, login string) (*entities_account.Account, error)
	}
	ISelectAccountByEmailCommand interface {
		Execute(context context.Context, email string) (*entities_account.Account, error)
	}
	IUpdateAccountByIdCommand interface {
		Execute(context context.Context, id int, newAccount *entities_account.Account) error
	}
)

type (
	ISelectSessionByAccessTokenCommand interface {
		Execute(ctx context.Context, token string) (*session_entities.Session, error)
	}
	ISelectSessionsByAccountIdCommand interface {
		Execute(ctx context.Context, id int) ([]*session_entities.Session, error)
	}
	IInsertSessionCommand interface {
		Execute(ctx context.Context, session *session_entities.Session) error
	}
	IUpdateSessionByAccessTokenCommand interface {
		Execute(ctx context.Context, accessToken string, newSession *session_entities.Session) error
	}
	IDeleteSessionByAccessTokenCommand interface {
		Execute(ctx context.Context, accessToken string) error
	}
)

type (
	ICreateVerificationCommand interface {
		Execute(context context.Context, verification *verification_entities.Verification) error
	}
	ISelectVerificationsByAccountIdCommand interface {
		Execute(context context.Context, accountId int) ([]*verification_entities.Verification, error)
	}
	IDeleteVerificationsByAccountIdCommand interface {
		Execute(context context.Context, accountId int) error
	}
)

type (
	IInsertDeviceCommand interface {
		Execute(context context.Context, device *session_entities.Device) error
	}
	ISelectDeviceByIdCommand interface {
		Execute(context context.Context, id int) (*session_entities.Device, error)
	}
	ISelectDeviceByAccessTokenCommand interface {
		Execute(context context.Context, accessToken string) (*session_entities.Device, error)
	}
	ISelectDevicesByAccountIdCommand interface {
		Execute(context context.Context, accountId int) ([]*session_entities.Device, error)
	}
	IDeleteDeviceByIdCommand interface {
		Execute(context context.Context, id int) error
	}
	IUpdateDeviceByAccessTokenCommand interface {
		Execute(context context.Context, token string, newDevice *session_entities.Device) error
	}
)
