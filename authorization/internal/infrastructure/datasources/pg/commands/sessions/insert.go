package sessions

import (
	"authorization/internal/entities"
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

func (c *insertSessionPGCommand) Execute(ctx context.Context, session *entities.Session) (int, error) {
	sql, args, err := query_builders.NewInsertSessionQuery(&c.client.Builder, session)
	if err != nil {
		return 0, err
	}

	var id int
	err = c.client.Pool.QueryRow(ctx, sql, args...).Scan(&id)
	return id, err
}
