package avatars

import (
	"context"
	"smartri_app/internal/entities/avatar"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type insertAvatarPGCommand struct {
	client *postgres.Client
}

func NewInsertAvatarPGCommand(client *postgres.Client) datasources.IInsertAvatarCommand {
	return &insertAvatarPGCommand{client: client}
}

func (c *insertAvatarPGCommand) Execute(context context.Context, avatar *avatar.Avatar) error {
	sql, args, err := query_builders.NewInsertAvatarQuery(&c.client.Builder, avatar)
	if err != nil {
		return err
	}

	_, err = c.client.Pool.Exec(context, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
