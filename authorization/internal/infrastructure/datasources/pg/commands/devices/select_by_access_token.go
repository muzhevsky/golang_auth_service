package devices

import (
	"authorization/internal/entities/session_entities"
	"authorization/internal/infrastructure/datasources"
	"authorization/internal/infrastructure/datasources/pg/query_builders"
	"authorization/pkg/postgres"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

type selectDeviceByAccessTokenPGCommand struct {
	client *postgres.Client
}

func NewSelectDeviceByAccessTokenPGCommand(client *postgres.Client) datasources.ISelectDeviceByAccessTokenCommand {
	return &selectDeviceByAccessTokenPGCommand{client: client}
}

func (c *selectDeviceByAccessTokenPGCommand) Execute(context context.Context, accessToken string) (*session_entities.Device, error) {
	sql, args, err := query_builders.NewSelectDeviceByAccessTokenQuery(&c.client.Builder, accessToken)
	if err != nil {
		return nil, err
	}

	row := c.client.Pool.QueryRow(context, sql, args...)
	result := session_entities.Device{}
	err = row.Scan(&result.Id, &result.AccountId, &result.Name, &result.SessionAccessToken, &result.SessionCreationTime)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &result, err
}
