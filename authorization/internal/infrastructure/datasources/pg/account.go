package pg

import (
	"authorization/internal/entities"
	"authorization/pkg/postgres"
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
)

const accountTableName = "accounts"

type AccountDatasource struct {
	pg *postgres.Client
}

func NewAccountDatasource(pg *postgres.Client) *AccountDatasource {
	return &AccountDatasource{pg: pg}
}

func (ds *AccountDatasource) Create(context context.Context, account *entities.Account) (id int, err error) {
	sql, args, err := ds.pg.Builder.
		Insert(accountTableName).
		Columns("login", "password", "email", "nickname", "registration_date").
		Values(account.Login, account.Password, account.EMail, account.Nickname, account.CreationTime).
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

func (ds *AccountDatasource) SelectById(context context.Context, id int) (*entities.Account, error) {
	return ds.findOne(context, "id", id)
}

func (ds *AccountDatasource) SelectByLogin(context context.Context, login string) (*entities.Account, error) {
	return ds.findOne(context, "login", login)
}

func (ds *AccountDatasource) SelectByEmail(context context.Context, email string) (*entities.Account, error) {
	return ds.findOne(context, "email", email)
}

func (ds *AccountDatasource) findOne(context context.Context, columnName string, value interface{}) (result *entities.Account, err error) {
	sql, _, err := ds.pg.Builder.Select("id", "login", "password", "nickname", "email", "registration_date", "is_verified").
		From(accountTableName).
		Where(sq.Eq{columnName: value}).
		ToSql()
	if err != nil {
		return nil, err
	}

	result = &entities.Account{}
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

func (ds *AccountDatasource) UpdateById(context context.Context, id int, updateFunc func(*entities.Account)) error {
	user, err := ds.SelectById(context, id)
	if err != nil {
		return err
	}

	updateFunc(user)

	sql, args, err := ds.pg.Builder.Update(accountTableName).
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

func (ds *AccountDatasource) DeleteById(context context.Context, id int) error {
	sql, _, err := ds.pg.Builder.Delete(accountTableName).Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}

	_, err = ds.pg.Pool.Exec(context, sql)
	return err
}
