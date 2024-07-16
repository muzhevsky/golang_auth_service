package pg

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func WrapError(context context.Context, err error, tx pgx.Tx) error {
	if err != nil {
		tx.Rollback(context)
	}

	return err
}
