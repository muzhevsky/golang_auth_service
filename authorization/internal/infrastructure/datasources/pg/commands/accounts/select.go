package accounts

import (
	"authorization/internal/entities/account"
	"authorization/pkg/postgres"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

func selectAccount(context context.Context, client *postgres.Client, sql string, args []any) (*account.Account, error) {
	result := &account.Account{}
	row := client.Pool.QueryRow(context, sql, args...)
	err := row.Scan(&result.Id, &result.Login, &result.Password, &result.Nickname, &result.Email, &result.RegistrationDate, &result.IsVerified)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}
