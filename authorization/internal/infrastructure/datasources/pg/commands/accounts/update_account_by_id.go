package accounts

import (
	"authorization/internal/entities/account"
	"authorization/internal/infrastructure/datasources/pg/query_builders"
	"authorization/pkg/postgres"
	"context"
)

type updateAccountPGCommand struct {
	client *postgres.Client
}

func NewUpdateAccountByIdPGCommand(client *postgres.Client) *updateAccountPGCommand {
	return &updateAccountPGCommand{client: client}
}

func (c *updateAccountPGCommand) Execute(context context.Context, id int, newAccount *account.Account) error {
	sql, args, err := query_builders.NewUpdateAccountByIdQuery(&c.client.Builder, id, newAccount)
	if err != nil {
		return err
	}

	_, err = c.client.Pool.Exec(context, sql, args...)
	return err
}
