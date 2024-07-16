package devices

import (
	"authorization/internal/infrastructure/datasources"
	"authorization/internal/infrastructure/datasources/pg/query_builders"
	"authorization/pkg/postgres"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type deleteDeviceByIdPGCommand struct {
	client *postgres.Client
}

func NewDeleteDeviceByIdPGCommand(client *postgres.Client) datasources.IDeleteDeviceByIdCommand {
	return &deleteDeviceByIdPGCommand{client: client}
}

func (c *deleteDeviceByIdPGCommand) Execute(context context.Context, id int) error {
	sql, args, err := query_builders.NewDeleteDeviceByIdQuery(&c.client.Builder, id)
	if err != nil {
		return err
	}

	_, err = c.client.Pool.Exec(context, sql, args...)
	fmt.Println(err)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil
	}

	return err
}
