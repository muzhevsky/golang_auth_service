package sessions

import (
	"authorization/internal/entities/session_entities"
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

func (s *selectSessionByAccountIdPGCommand) Execute(ctx context.Context, id int) ([]*session_entities.Session, error) {
	sql, args, err := query_builders.NewSelectSessionsByAccountIdQuery(&s.client.Builder, id)

	if err != nil {
		return nil, err
	}

	rows, err := s.client.Pool.Query(ctx, sql, args...)

	if err != nil {
		return nil, err
	}

	result := make([]*session_entities.Session, 0)

	for rows.Next() {
		row := session_entities.Session{}
		err = rows.Scan(&row.AccessToken, &row.RefreshToken, &row.AccountId, &row.ExpiresAt)
		if err != nil {
			return nil, err
		}
		result = append(result, &row)
	}

	return result, err
}
