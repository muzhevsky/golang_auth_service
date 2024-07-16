package sessions

import (
	"authorization/internal/entities/session_entities"
	"authorization/internal/infrastructure/datasources"
	"authorization/internal/infrastructure/datasources/pg/query_builders"
	"authorization/pkg/postgres"
	"context"
)

type insertSessionPGCommand struct {
	client *postgres.Client
}

func NewInsertSessionPGCommand(client *postgres.Client) datasources.IInsertSessionCommand {
	return &insertSessionPGCommand{client: client}
}

func (c *insertSessionPGCommand) Execute(ctx context.Context, session *session_entities.Session) error {
	sql, args, err := query_builders.NewInsertSessionQuery(&c.client.Builder, session)
	if err != nil {
		return err
	}

	var id int
	err = c.client.Pool.QueryRow(ctx, sql, args...).Scan(&id)
	return err
}
