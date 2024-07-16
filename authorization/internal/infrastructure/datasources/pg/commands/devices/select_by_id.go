package devices

import (
	"authorization/internal/entities/session"
	"authorization/internal/infrastructure/datasources"
	"authorization/internal/infrastructure/datasources/pg/query_builders"
	"authorization/pkg/postgres"
	"context"
)

type selectDeviceByIdPGCommand struct {
	client *postgres.Client
}

func NewSelectDeviceByIdPGCommand(client *postgres.Client) datasources.ISelectDeviceByIdCommand {
	return &selectDeviceByIdPGCommand{client: client}
}

func (c *selectDeviceByIdPGCommand) Execute(context context.Context, id int) (*session.Device, error) {
	sql, args, err := query_builders.NewSelectDeviceByIdQuery(&c.client.Builder, id)
	if err != nil {
		return nil, err
	}

	row := c.client.Pool.QueryRow(context, sql, args...)
	device := session.Device{}
	err = row.Scan(&device.Id, &device.AccountId, &device.Name, &device.SessionAccessToken, &device.SessionCreationTime)
	if err != nil {
		return nil, err
	}

	return &device, nil
}
