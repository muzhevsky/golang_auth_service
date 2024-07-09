package datasources

import (
	"authorization/internal/entities/account"
	"authorization/internal/entities/session"
	"authorization/internal/entities/verification"
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
	ISelectSessionByAccessTokenCommand interface {
		Execute(ctx context.Context, token string) (*session.Session, error)
	}
	ISelectSessionsByAccountIdCommand interface {
		Execute(ctx context.Context, id int) ([]*session.Session, error)
	}
	IInsertSessionCommand interface {
		Execute(ctx context.Context, session *session.Session) error
	}
	IUpdateSessionByAccessTokenCommand interface {
		Execute(ctx context.Context, accessToken string, newSession *session.Session) error
	}
)

type (
	ICreateVerificationCommand interface {
		Execute(context context.Context, verification *verification.Verification) error
	}
	ISelectVerificationsByAccountIdCommand interface {
		Execute(context context.Context, accountId int) ([]*verification.Verification, error)
	}
	IDeleteVerificationsByAccountIdCommand interface {
		Execute(context context.Context, accountId int) error
	}
)
