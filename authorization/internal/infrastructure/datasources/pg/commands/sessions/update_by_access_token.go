package sessions

import (
	"authorization/internal/entities/session_entities"
	"authorization/internal/infrastructure/datasources"
	"authorization/internal/infrastructure/datasources/pg/query_builders"
	"authorization/pkg/postgres"
	"context"
)

type updateSessionByAccessTokenPGCommand struct {
	client *postgres.Client
}

func NewUpdateSessionByIdPGCommand(client *postgres.Client) datasources.IUpdateSessionByAccessTokenCommand {
	return &updateSessionByAccessTokenPGCommand{client: client}
}

func (s *updateSessionByAccessTokenPGCommand) Execute(ctx context.Context, accessToken string, session *session_entities.Session) error {
	sql, args, err := query_builders.NewUpdateSessionQuery(&s.client.Builder, accessToken, session)
	if err != nil {
		return err
	}

	_, err = s.client.Pool.Exec(ctx, sql, args...)
	return err
}
