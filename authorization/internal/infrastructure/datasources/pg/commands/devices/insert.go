package devices

import (
	"authorization/internal/entities/session_entities"
	"authorization/internal/infrastructure/datasources"
	"authorization/internal/infrastructure/datasources/pg/query_builders"
	"authorization/pkg/postgres"
	"context"
)

type insertDevicePGCommand struct {
	client *postgres.Client
}

func NewInsertDevicePGCommand(client *postgres.Client) datasources.IInsertDeviceCommand {
	return &insertDevicePGCommand{client: client}
}

func (c *insertDevicePGCommand) Execute(context context.Context, device *session_entities.Device) error {
	sql, args, err := query_builders.NewInsertDeviceQuery(&c.client.Builder, device)
	if err != nil {
		return err
	}

	_, err = c.client.Pool.Exec(context, sql, args...)
	return err
}
