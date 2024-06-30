package datasources

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type ITransaction interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type pgTransaction struct {
	transaction pgx.Tx
}

func NewPgTransaction(transaction pgx.Tx) ITransaction {
	return &pgTransaction{transaction: transaction}
}

func (p *pgTransaction) Commit(ctx context.Context) error {
	err := p.transaction.Commit(ctx)
	return err
}

func (p *pgTransaction) Rollback(ctx context.Context) error {
	err := p.transaction.Rollback(ctx)
	return err
}
