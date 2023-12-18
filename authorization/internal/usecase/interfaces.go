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

	IPasswordHashProvider interface {
		GenerateHashPassword(stringToHash string) ([]byte, error)
		CompareStringAndHash(stringToCompare string, hashedString string) bool
	}

	IUserRepo interface {
		Create(context context.Context, user *entities.User) (int, error)
		FindOne(context context.Context, user *entities.User) (*entities.User, error)
		CheckLoginExist(context context.Context, login string) (bool, error)
		CheckEmailExist(context context.Context, email string) (bool, error)
		Verify(context context.Context, id int) error
	}

	IVerification interface {
		CreateVerification(user *entities.User) error
		Verify(context context.Context, verification *entities.Verification) (bool, error)
	}

	IVerificationRepo interface {
		Create(verification *entities.Verification) error
		FindOne(userId int) (*entities.Verification, error)
	}

	ISession interface {
		VerifyAccessToken(token string) (bool, error)
		CreateTokens(user *entities.User) (*entities.Session, error)
		UpdateAccessToken(accessToken, refreshToken string) (*entities.Session, error)
	}

	ITokenManager interface {
		GenerateToken(claims map[string]interface{}) (string, error)
		ParseToken(token string) (map[string]interface{}, error)
	}

	ISessionRepo interface {
		Create(ctx context.Context, session *entities.Session) error
		Update(ctx context.Context, session *entities.Session) error
		FindByAccess(ctx context.Context, token string) (*entities.Session, error)
	}

	IMailer interface {
		SendMail(receiver string, subject string, body string)
	}

	IUserDataRepo interface {
	}
)
