package user_data

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"smartri_app/internal/entities"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectUserDataByAccountId struct {
	client *postgres.Client
}

func NewSelectUserDataByAccountId(client *postgres.Client) datasources.ISelectUserDataByAccountIdCommand {
	return &selectUserDataByAccountId{client: client}
}

func (u *selectUserDataByAccountId) Execute(context context.Context, accountId int) (*entities.UserData, error) {
	sql, args, err := query_builders.NewSelectUserDataByAccountIdQuery(&u.client.Builder, accountId)
	if err != nil {
		return nil, err
	}

	row := u.client.Pool.QueryRow(context, sql, args...)
	var result entities.UserData
	err = row.Scan(&result.Age, &result.Gender, &result.XP, &result.AccountId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
