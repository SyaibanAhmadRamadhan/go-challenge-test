package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UOWRepository interface {
	StartTx(ctx context.Context, opts pgx.TxOptions, fn func() error) error
	GetTx() (pgx.Tx, error)
	GetDB() (*pgxpool.Pool, error)
}

func LevelReadCommitted() pgx.TxOptions {
	return pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadWrite,
		DeferrableMode: "",
		BeginQuery:     "",
	}
}
