package accounts

import (
	"authorization/internal/entities/entities_account"
	"authorization/pkg/postgres"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

func selectAccount(context context.Context, client *postgres.Client, sql string, args []any) (*entities_account.Account, error) {
	result := &entities_account.Account{}
	row := client.Pool.QueryRow(context, sql, args...)
	err := row.Scan(&result.Id, &result.Login, &result.Password, &result.Email, &result.RegistrationDate, &result.IsVerified)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}
