package pg

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"smartri_app/internal/entities"
	"smartri_app/pkg/postgres"
)

type usersDataSource struct {
	client *postgres.Client
}

func NewUsersDataSource(client *postgres.Client) *usersDataSource {
	return &usersDataSource{client: client}
}

func (u *usersDataSource) SelectByAccountId(context context.Context, accountId int) (*entities.User, error) {
	sql, args, err := u.client.Builder.
		Select("age", "gender", "XP", "account_id").
		From("user_details").
		Where(squirrel.Eq{"account_id": accountId}).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := u.client.Pool.QueryRow(context, sql, args...)
	var result entities.User
	err = row.Scan(&result.Age, &result.Gender, &result.XP, &result.AccountId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

func (u *usersDataSource) Insert(context context.Context, user *entities.User) error {
	sql, args, err := u.client.Builder.
		Insert("user_details").
		Columns("age", "gender", "xp", "account_id").
		Values(user.Age, user.Gender, 0, user.AccountId).
		ToSql()
	if err != nil {
		return err
	}

	_, err = u.client.Pool.Exec(context, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (u *usersDataSource) UpdateByAccountId(context context.Context, accountId int, details *entities.User) error {
	user, err := u.SelectByAccountId(context, accountId)
	if err != nil || user == nil {
		return err
	}

	sql, args, err := u.client.Builder.Update("user_details").
		Set("xp", details.XP).
		Set("age", details.Age).
		Set("gender", details.Gender).
		Where(squirrel.Eq{"account_id": user.AccountId}).ToSql()

	_, err = u.client.Pool.Exec(context, sql, args...)

	return err
}
