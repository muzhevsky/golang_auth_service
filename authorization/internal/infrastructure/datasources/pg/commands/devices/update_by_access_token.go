package devices

import (
	"authorization/internal/entities/session_entities"
	"authorization/internal/infrastructure/datasources"
	"authorization/internal/infrastructure/datasources/pg/query_builders"
	"authorization/pkg/postgres"
	"context"
)

type updateDeviceByAccessTokenPGCommand struct {
	client *postgres.Client
}

func NewUpdateDeviceByAccessTokenPGCommand(client *postgres.Client) datasources.IUpdateDeviceByAccessTokenCommand {
	return &updateDeviceByAccessTokenPGCommand{client: client}
}

func (c *updateDeviceByAccessTokenPGCommand) Execute(context context.Context, token string, newDevice *session_entities.Device) error {
	sql, args, err := query_builders.NewUpdateDeviceByAccessTokenQuery(&c.client.Builder, token, newDevice)
	if err != nil {
		return err
	}

	_, err = c.client.Pool.Exec(context, sql, args...)
	return err
}
