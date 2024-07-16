package accounts

import (
	"authorization/internal/entities/entities_account"
	"authorization/internal/infrastructure/datasources/pg/query_builders"
	"authorization/pkg/postgres"
	"context"
)

type selectAccountByEmailCommand struct {
	client *postgres.Client
}

func NewSelectAccountByEmailPGCommand(client *postgres.Client) *selectAccountByEmailCommand {
	return &selectAccountByEmailCommand{client: client}
}

func (s *selectAccountByEmailCommand) Execute(context context.Context, email string) (*entities_account.Account, error) {
	sql, args, err := query_builders.NewSelectAccountByEmailQuery(&s.client.Builder, email)
	if err != nil {
		return nil, err
	}

	return selectAccount(context, s.client, sql, args)
}
