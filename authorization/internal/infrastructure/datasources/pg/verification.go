package pg

import (
	"authorization/internal/entities"
	"authorization/pkg/postgres"
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
)

const verificationTableName = "verification_codes"

type pgVerificationDatasource struct {
	pg *postgres.Client
}

func NewPgVerificationDatasource(pg *postgres.Client) *pgVerificationDatasource {
	return &pgVerificationDatasource{pg: pg}
}

func (ds *pgVerificationDatasource) Create(context context.Context, verification *entities.Verification) (id int, err error) {
	sql, args, err := ds.pg.Builder.
		Insert(verificationTableName).
		Columns("user_id", "code", "expiration_time").
		Values(verification.UserId, verification.Code, verification.ExpirationTime).
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

func (ds *pgVerificationDatasource) SelectById(context context.Context, id int) (*entities.Verification, error) {
	sql, args, err := ds.pg.Builder.Select("id", "user_id", "code", "expiration_time").
		From(verificationTableName).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	result := &entities.Verification{}
	row := ds.pg.Pool.QueryRow(context, sql, args...)
	err = row.Scan(&result.Id, &result.UserId, &result.Code, &result.ExpirationTime)
	if err != nil {
		if errors.Is(err, ds.pg.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

func (ds *pgVerificationDatasource) SelectByUserId(context context.Context, userId int) ([]*entities.Verification, error) {
	sql, args, err := ds.pg.Builder.Select("id", "user_id", "code", "expiration_time").
		From(verificationTableName).
		Where(sq.Eq{"user_id": userId}).
		ToSql()
	if err != nil {
		return nil, err
	}

	result := make([]*entities.Verification, 0)
	rows, err := ds.pg.Pool.Query(context, sql, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		verification := &entities.Verification{}
		err := rows.Scan(&verification.Id, &verification.UserId, &verification.Code, &verification.ExpirationTime)
		if err != nil {
			return nil, err
		}
		result = append(result, verification)
	}

	return result, nil
}

func (ds *pgVerificationDatasource) DeleteById(context context.Context, id int) error {
	sql, args, err := ds.pg.Builder.Delete(verificationTableName).Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}

	_, err = ds.pg.Pool.Exec(context, sql, args...)
	return err
}
