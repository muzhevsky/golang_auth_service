package accounts

import (
	"authorization/internal/entities/account"
	"authorization/internal/infrastructure/datasources/pg/query_builders"
	"authorization/pkg/postgres"
	"context"
)

type selectAccountByLoginCommand struct {
	client *postgres.Client
}

func NewSelectAccountByLoginPGCommand(client *postgres.Client) *selectAccountByLoginCommand {
	return &selectAccountByLoginCommand{client: client}
}

func (s *selectAccountByLoginCommand) Execute(context context.Context, login string) (*account.Account, error) {
	sql, args, err := query_builders.NewSelectAccountByLoginQuery(&s.client.Builder, login)
	if err != nil {
		return nil, err
	}

	return selectAccount(context, s.client, sql, args)
}
