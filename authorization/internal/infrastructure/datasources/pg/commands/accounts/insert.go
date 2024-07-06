package accounts

import (
	"authorization/internal/entities/account"
	"authorization/internal/infrastructure/datasources/pg/query_builders"
	"authorization/pkg/postgres"
	"context"
)

type insertAccountPGCommand struct {
	client *postgres.Client
}

func NewInsertAccountPGCommand(client *postgres.Client) *insertAccountPGCommand {
	return &insertAccountPGCommand{client: client}
}

func (c *insertAccountPGCommand) Execute(context context.Context, account *account.Account) (int, error) {
	sql, args, err := query_builders.NewInsertAccountQuery(&c.client.Builder, account)
	if err != nil {
		return 0, err
	}

	id := 0
	err = c.client.Pool.QueryRow(context, sql, args...).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
