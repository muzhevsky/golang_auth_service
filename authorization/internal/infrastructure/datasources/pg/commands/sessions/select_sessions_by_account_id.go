package sessions

import (
	"authorization/internal/entities"
	"authorization/internal/infrastructure/datasources"
	"authorization/internal/infrastructure/datasources/pg/query_builders"
	"authorization/pkg/postgres"
	"context"
)

type selectSessionByAccountIdPGCommand struct {
	client *postgres.Client
}

func NewSelectSessionByAccountIdPGCommand(client *postgres.Client) datasources.ISelectSessionsByAccountIdCommand {
	return &selectSessionByAccountIdPGCommand{client: client}
}

func (s *selectSessionByAccountIdPGCommand) Execute(ctx context.Context, id int) ([]*entities.Session, error) {
	sql, args, err := query_builders.NewSelectSessionsByAccountIdQuery(&s.client.Builder, id)

	if err != nil {
		return nil, err
	}

	rows, err := s.client.Pool.Query(ctx, sql, args...)

	if err != nil {
		return nil, err
	}

	result := make([]*entities.Session, 0)

	for rows.Next() {
		row := entities.Session{}
		err = rows.Scan(&row.Id, &row.AccessToken, &row.RefreshToken, &row.AccountId, &row.ExpiresAt)
		if err != nil {
			return nil, err
		}
		result = append(result, &row)
	}

	return result, err
}
