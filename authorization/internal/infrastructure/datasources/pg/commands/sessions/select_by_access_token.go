package sessions

import (
	"authorization/internal/entities/session_entities"
	"authorization/internal/infrastructure/datasources"
	"authorization/internal/infrastructure/datasources/pg/query_builders"
	"authorization/pkg/postgres"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

type selectSessionByAccessTokenPGCommand struct {
	client *postgres.Client
}

func NewSelectSessionByAccessTokenPGCommand(client *postgres.Client) datasources.ISelectSessionByAccessTokenCommand {
	return &selectSessionByAccessTokenPGCommand{client: client}
}

func (s *selectSessionByAccessTokenPGCommand) Execute(ctx context.Context, token string) (*session_entities.Session, error) {
	sql, args, err := query_builders.NewSelectSessionByAccessTokenQuery(&s.client.Builder, token)

	if err != nil {
		return nil, err
	}

	result := &session_entities.Session{}
	err = s.client.Pool.
		QueryRow(ctx, sql, args...).
		Scan(&result.AccessToken, &result.RefreshToken, &result.AccountId, &result.ExpiresAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	return result, err
}
