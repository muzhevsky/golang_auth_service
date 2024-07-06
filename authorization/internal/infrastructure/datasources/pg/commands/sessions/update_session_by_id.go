package sessions

import (
	"authorization/internal/entities"
	"authorization/internal/infrastructure/datasources"
	"authorization/internal/infrastructure/datasources/pg/query_builders"
	"authorization/pkg/postgres"
	"context"
)

type updateSessionByIdPGCommand struct {
	client *postgres.Client
}

func NewUpdateSessionByIdPGCommand(client *postgres.Client) datasources.IUpdateSessionByIdCommand {
	return &updateSessionByIdPGCommand{client: client}
}

func (s *updateSessionByIdPGCommand) Execute(ctx context.Context, session *entities.Session) error {
	sql, args, err := query_builders.NewUpdateSessionQuery(&s.client.Builder, session)
	if err != nil {
		return err
	}

	_, err = s.client.Pool.Exec(ctx, sql, args...)
	return err
}
