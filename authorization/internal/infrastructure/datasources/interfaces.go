package datasources

import (
	"authorization/internal/entities"
	"context"
)

type (
	IUserDataSource interface {
		Create(context context.Context, user *entities.User) (int, error)
		SelectById(context context.Context, id int) (*entities.User, error)
		SelectByLogin(context context.Context, login string) (*entities.User, error)
		SelectByEmail(context context.Context, email string) (*entities.User, error)
		UpdateById(context context.Context, id int, updateFunc func(*entities.User)) error
		DeleteById(context context.Context, id int) error
	}

	IVerificationDataSource interface {
		Create(context context.Context, user *entities.Verification) (int, error)
		SelectById(context context.Context, id int) (*entities.Verification, error)
		SelectByUserId(context context.Context, userId int) ([]*entities.Verification, error)
		DeleteById(context context.Context, id int) error
	}

	ISessionDataSource interface {
		Create(ctx context.Context, session *entities.Session) (int, error)
		SelectByAccess(ctx context.Context, token string) (*entities.Session, error)
		SelectByUserId(ctx context.Context, userId int) ([]*entities.Session, error)
		Delete(ctx context.Context, session *entities.Session) error
	}
)
