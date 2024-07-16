package accounts

import (
	"authorization/internal/entities/entities_account"
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

func (s *selectAccountByIdCommand) Execute(context context.Context, id int) (*entities_account.Account, error) {
	sql, args, err := query_builders.NewSelectAccountByIdQuery(&s.client.Builder, id)
	if err != nil {
		return nil, err
	}

	return selectAccount(context, s.client, sql, args)
}
