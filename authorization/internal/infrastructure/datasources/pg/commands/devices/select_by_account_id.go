package devices

import (
	"authorization/internal/entities/session"
	"authorization/internal/infrastructure/datasources"
	"authorization/internal/infrastructure/datasources/pg/query_builders"
	"authorization/pkg/postgres"
	"context"
)

type selectDevicesByAccountIdPGCommand struct {
	client *postgres.Client
}

func NewSelectDevicesByAccountIdPGCommand(client *postgres.Client) datasources.ISelectDevicesByAccountIdCommand {
	return &selectDevicesByAccountIdPGCommand{client: client}
}

func (c *selectDevicesByAccountIdPGCommand) Execute(context context.Context, accountId int) ([]*session.Device, error) {
	sql, args, err := query_builders.NewSelectDeviceByAccountIdQuery(&c.client.Builder, accountId)
	if err != nil {
		return nil, err
	}

	rows, err := c.client.Pool.Query(context, sql, args...)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	result := make([]*session.Device, 0)
	for rows.Next() {
		curr := session.Device{}
		err = rows.Scan(&curr.Id, &curr.AccountId, &curr.Name, &curr.SessionAccessToken, &curr.SessionCreationTime)
		if err != nil {
			return nil, err
		}
		result = append(result, &curr)
	}

	return result, err
}
