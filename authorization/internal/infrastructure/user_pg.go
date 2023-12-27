package infrastructure

import (
	"authorization/internal/entities"
	"authorization/pkg/postgres"
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type userRepo struct {
	pg *postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *userRepo {
	return &userRepo{pg}
}

func (u *userRepo) Create(context context.Context, user *entities.User) (id int, err error) {
	sql, args, err := u.pg.Builder.
		Insert("users").
		Columns("login", "password", "email", "nickname", "registration_date").
		Values(user.Login, user.Password, user.EMail, user.Nickname, user.CreationTime).
		Suffix("RETURNING \"id\"").
		ToSql()
	if err != nil {
		return 0, err
	}

	err = u.pg.Pool.QueryRow(context, sql, args...).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *userRepo) FindById(context context.Context, id int) (*entities.User, error) {
	return u.findOne(context, "id", id)
}

func (u *userRepo) FindByLogin(context context.Context, login string) (*entities.User, error) {
	return u.findOne(context, "login", login)
}

func (u *userRepo) FindByEmail(context context.Context, email string) (*entities.User, error) {
	return u.findOne(context, "email", email)
}

func (u *userRepo) findOne(context context.Context, columnName string, value interface{}) (result *entities.User, err error) {
	sql, _, err := u.pg.Builder.Select("id", "login", "password", "nickname", "email", "registration_date", "is_verified").
		From("users").
		Where(sq.Eq{columnName: value}).
		ToSql()
	if err != nil {
		return nil, err
	}

	result = &entities.User{}
	err = u.pg.Pool.QueryRow(context, sql, value).
		Scan(&result.Id, &result.Login, &result.Password, &result.Nickname, &result.EMail, &result.CreationTime, &result.IsVerified)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entities.UserNotFound
	}
	return result, nil
}

func (u *userRepo) CheckLoginExist(context context.Context, login string) (result bool, err error) {
	sql, args, err := u.pg.Builder.
		Select("login").From("users").
		Where(sq.Eq{"login": login}).Limit(1).
		ToSql()
	if err != nil {
		return false, err
	}

	err = u.pg.Pool.QueryRow(context, sql, args...).Scan(&login)
	if err != nil {
		if errors.Is(err, u.pg.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, err
}

func (u *userRepo) CheckEmailExist(context context.Context, email string) (bool, error) {
	sql, args, err := u.pg.Builder.
		Select("email").From("users").
		Where(sq.Eq{"email": email}).Limit(1).
		ToSql()
	if err != nil {
		return false, err
	}

	err = u.pg.Pool.QueryRow(context, sql, args...).Scan(&email)
	if err != nil {
		if err == u.pg.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, err
}

func (u *userRepo) Verify(context context.Context, id int) error {
	request, args, err := u.pg.Builder.
		Update("users").Set("is_verified", true).Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}

	_, err = u.pg.Pool.Query(context, request, args...)
	return err
}
