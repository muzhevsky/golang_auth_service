package user_data

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"smartri_app/internal/entities/user_data"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectUserDataByAccountIdPGCommand struct {
	client *postgres.Client
}

func NewSelectUserDataByAccountIdPGCommand(client *postgres.Client) datasources.ISelectUserDataByAccountIdCommand {
	return &selectUserDataByAccountIdPGCommand{client: client}
}

func (u *selectUserDataByAccountIdPGCommand) Execute(context context.Context, accountId int) (*user_data.UserData, error) {
	sql, args, err := query_builders.NewSelectUserDataByAccountIdQuery(&u.client.Builder, accountId)
	if err != nil {
		return nil, err
	}

	row := u.client.Pool.QueryRow(context, sql, args...)
	var result user_data.UserData
	err = row.Scan(&result.Age, &result.Gender, &result.XP, &result.AccountId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
