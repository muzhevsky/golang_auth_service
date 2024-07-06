package datasources

import (
	"authorization/internal/entities"
	"authorization/internal/entities/account"
	"context"
)

type (
	IInsertAccountCommand interface {
		Execute(context context.Context, user *account.Account) (int, error)
	}
	ISelectAccountByIdCommand interface {
		Execute(context context.Context, id int) (*account.Account, error)
	}
	ISelectAccountByLoginCommand interface {
		Execute(context context.Context, login string) (*account.Account, error)
	}
	ISelectAccountByEmailCommand interface {
		Execute(context context.Context, email string) (*account.Account, error)
	}
	IUpdateAccountByIdCommand interface {
		Execute(context context.Context, id int, newAccount *account.Account) error
	}
)

type (
	ISelectSessionByIdCommand interface {
		Execute(ctx context.Context, id int) (*entities.Session, error)
	}
	ISelectSessionByAccessTokenCommand interface {
		Execute(ctx context.Context, token string) (*entities.Session, error)
	}
	ISelectSessionsByAccountIdCommand interface {
		Execute(ctx context.Context, id int) ([]*entities.Session, error)
	}
	IInsertSessionCommand interface {
		Execute(ctx context.Context, session *entities.Session) (int, error)
	}
	IUpdateSessionByIdCommand interface {
		Execute(ctx context.Context, session *entities.Session) error
	}
)

type (
	IVerificationDataSource interface {
		Create(context context.Context, user *entities.Verification) (int, error)
		SelectById(context context.Context, id int) (*entities.Verification, error)
		SelectByUserId(context context.Context, userId int) ([]*entities.Verification, error)
		DeleteById(context context.Context, id int) error
	}
)
