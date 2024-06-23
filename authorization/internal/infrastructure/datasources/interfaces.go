package datasources

import (
	"authorization/internal/entities"
	"context"
)

type (
	IAccountDataSource interface {
		Create(context context.Context, user *entities.Account) (int, error)
		SelectById(context context.Context, id int) (*entities.Account, error)
		SelectByLogin(context context.Context, login string) (*entities.Account, error)
		SelectByEmail(context context.Context, email string) (*entities.Account, error)
		UpdateById(context context.Context, id int, updateFunc func(*entities.Account)) error
		DeleteById(context context.Context, id int) error
	}

	IVerificationDataSource interface {
		Create(context context.Context, user *entities.Verification) (int, error)
		SelectById(context context.Context, id int) (*entities.Verification, error)
		SelectByUserId(context context.Context, userId int) ([]*entities.Verification, error)
		DeleteById(context context.Context, id int) error
	}

	ISessionDatasource interface {
		Create(ctx context.Context, session *entities.Session) (int, error)
		SelectByAccess(ctx context.Context, token string) (*entities.Session, error)
		SelectByUserId(ctx context.Context, userId int) ([]*entities.Session, error)
		UpdateById(context context.Context, id int, session *entities.Session) error
		Delete(ctx context.Context, session *entities.Session) error
	}
)
