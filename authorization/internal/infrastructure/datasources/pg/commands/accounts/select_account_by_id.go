package accounts

import (
	"authorization/internal/entities/account"
	"authorization/internal/infrastructure/datasources/pg/query_builders"
	"authorization/pkg/postgres"
	"context"
)

type selectAccountByIdCommand struct {
	client *postgres.Client
}

func NewSelectAccountByIdPGCommand(client *postgres.Client) *selectAccountByIdCommand {
	return &selectAccountByIdCommand{client: client}
}

func (s *selectAccountByIdCommand) Execute(context context.Context, id int) (*account.Account, error) {
	sql, args, err := query_builders.NewSelectAccountByIdQuery(&s.client.Builder, id)
	if err != nil {
		return nil, err
	}

	return selectAccount(context, s.client, sql, args)
}
