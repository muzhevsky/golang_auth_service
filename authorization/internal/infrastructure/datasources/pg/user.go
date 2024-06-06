package pg

import (
	"authorization/internal/entities"
	"authorization/pkg/postgres"
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
)

const userTableName = "users"

type pgUserDatasource struct {
	pg *postgres.Postgres
}

func NewPgUserDatasource(pg *postgres.Postgres) *pgUserDatasource {
	return &pgUserDatasource{pg: pg}
}

func (ds *pgUserDatasource) Create(context context.Context, user *entities.User) (id int, err error) {
	sql, args, err := ds.pg.Builder.
		Insert(userTableName).
		Columns("login", "password", "email", "nickname", "registration_date").
		Values(user.Login, user.Password, user.EMail, user.Nickname, user.CreationTime).
		Suffix("RETURNING \"id\"").
		ToSql()
	if err != nil {
		return 0, err
	}

	err = ds.pg.Pool.QueryRow(context, sql, args...).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ds *pgUserDatasource) SelectById(context context.Context, id int) (*entities.User, error) {
	return ds.findOne(context, "id", id)
}

func (ds *pgUserDatasource) SelectByLogin(context context.Context, login string) (*entities.User, error) {
	return ds.findOne(context, "login", login)
}

func (ds *pgUserDatasource) SelectByEmail(context context.Context, email string) (*entities.User, error) {
	return ds.findOne(context, "email", email)
}

func (ds *pgUserDatasource) findOne(context context.Context, columnName string, value interface{}) (result *entities.User, err error) {
	sql, _, err := ds.pg.Builder.Select("id", "login", "password", "nickname", "email", "registration_date", "is_verified").
		From(userTableName).
		Where(sq.Eq{columnName: value}).
		ToSql()
	if err != nil {
		return nil, err
	}

	result = &entities.User{}
	row := ds.pg.Pool.QueryRow(context, sql, value)
	err = row.Scan(&result.Id, &result.Login, &result.Password, &result.Nickname, &result.EMail, &result.CreationTime, &result.IsVerified)
	if err != nil {
		if errors.Is(err, ds.pg.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

func (ds *pgUserDatasource) UpdateById(context context.Context, id int, updateFunc func(*entities.User)) error {
	user, err := ds.SelectById(context, id)
	if err != nil {
		return err
	}

	updateFunc(user)

	sql, args, err := ds.pg.Builder.Update(userTableName).
		Set("login", user.Login).
		Set("email", user.EMail).
		Set("nickname", user.Nickname).
		Set("password", user.Password).
		Set("is_verified", user.IsVerified).
		Set("registration_date", user.CreationTime).
		Where(sq.Eq{"id": user.Id}).ToSql()

	_, err = ds.pg.Pool.Exec(context, sql, args...)
	return err
}

func (ds *pgUserDatasource) DeleteById(context context.Context, id int) error {
	sql, _, err := ds.pg.Builder.Delete(userTableName).Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}

	_, err = ds.pg.Pool.Exec(context, sql)
	return err
}
