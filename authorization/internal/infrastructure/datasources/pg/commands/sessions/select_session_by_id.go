package sessions

import (
	"authorization/internal/entities"
	"authorization/internal/infrastructure/datasources"
	"authorization/internal/infrastructure/datasources/pg/query_builders"
	"authorization/pkg/postgres"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

type selectSessionByIdPGCommand struct {
	client *postgres.Client
}

func NewSelectSessionByIdPGCommand(client *postgres.Client) datasources.ISelectSessionByIdCommand {
	return &selectSessionByIdPGCommand{client: client}
}

func (s *selectSessionByIdPGCommand) Execute(ctx context.Context, id int) (*entities.Session, error) {
	sql, args, err := query_builders.NewSelectSessionByIdQuery(&s.client.Builder, id)

	if err != nil {
		return nil, err
	}

	result := &entities.Session{}
	err = s.client.Pool.
		QueryRow(ctx, sql, args...).
		Scan(&result.Id, &result.AccessToken, &result.RefreshToken, &result.AccountId, &result.ExpiresAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	return result, err
}
