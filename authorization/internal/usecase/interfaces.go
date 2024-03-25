package usecase

import (
	"authorization/internal/entities"
	"context"
)

type (
	IUser interface {
		CreateUser(context context.Context, user *entities.User) (*entities.User, error)
		SignIn(context context.Context, user *entities.User) (*entities.User, error)
	}

	IHashProvider interface {
		GenerateHash(stringToHash string) ([]byte, error)
		CompareStringAndHash(stringToCompare string, hashedString string) bool
	}

	IUserRepo interface {
		Create(context context.Context, user *entities.User) (int, error)
		FindById(context context.Context, id int) (*entities.User, error)
		FindByLogin(context context.Context, login string) (*entities.User, error)
		FindByEmail(context context.Context, email string) (*entities.User, error)
		CheckLoginExist(context context.Context, login string) (bool, error)
		CheckEmailExist(context context.Context, email string) (bool, error)
		Verify(context context.Context, id int) error
	}

	IVerification interface {
		CreateVerification(context context.Context, user *entities.User) error
		Verify(context context.Context, verification *entities.Verification) error
	}

	IVerificationRepo interface {
		Create(context context.Context, verification *entities.Verification) error
		FindOne(context context.Context, userId int) (*entities.Verification, error)
		Clear(context context.Context, userId int) error
	}

	ISession interface {
		VerifyAccessToken(context context.Context, token string) error
		GetSession(context context.Context, token string) (*entities.Session, error)
		CreateSession(context context.Context, user *entities.User) (*entities.Session, error)
		UpdateSession(context context.Context, session *entities.Session) (*entities.Session, error)
	}

	IAccessTokenManager interface {
		GenerateToken(claims map[string]interface{}) (string, error)
		ParseToken(token string) (map[string]interface{}, error)
	}

	IRefreshTokenGenerator interface {
		GenerateToken(userId int) (string, error)
	}

	ISessionRepo interface {
		Create(ctx context.Context, session *entities.Session) error
		Update(ctx context.Context, session *entities.Session) error
		Delete(ctx context.Context, session *entities.Session) error
		FindByAccess(ctx context.Context, token string) (*entities.Session, error)
	}

	ISecurity interface {
		CheckAccess(ctx context.Context, route string, userId int) (bool, error)
	}

	IMailer interface {
		SendMail(receiver string, subject string, body string)
	}

	IUserDataRepo interface {
	}
)
