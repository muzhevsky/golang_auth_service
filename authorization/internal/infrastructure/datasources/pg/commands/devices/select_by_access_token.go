package devices

import (
	"authorization/internal/entities/session"
	"authorization/internal/infrastructure/datasources"
	"authorization/internal/infrastructure/datasources/pg/query_builders"
	"authorization/pkg/postgres"
	"context"
)

type selectDeviceByAccessTokenPGCommand struct {
	client *postgres.Client
}

func NewSelectDeviceByAccessTokenPGCommand(client *postgres.Client) datasources.ISelectDeviceByAccessTokenCommand {
	return &selectDeviceByAccessTokenPGCommand{client: client}
}

func (c *selectDeviceByAccessTokenPGCommand) Execute(context context.Context, accessToken string) (*session.Device, error) {
	sql, args, err := query_builders.NewSelectDeviceByAccessTokenQuery(&c.client.Builder, accessToken)
	if err != nil {
		return nil, err
	}

	row := c.client.Pool.QueryRow(context, sql, args...)
	result := session.Device{}
	err = row.Scan(&result.Id, &result.AccountId, &result.Name, &result.SessionAccessToken, &result.SessionCreationTime)
	if err != nil {
		return nil, err
	}

	return &result, err
}
