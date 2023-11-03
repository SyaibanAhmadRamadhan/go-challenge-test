package infra

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"challenge-test-synapsis/conf"
	"challenge-test-synapsis/helper"
)

func OpenConnectionDB(conf conf.PostgresConf) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	conn, err := pgxpool.New(ctx, conf.DBUrl())
	helper.PanicIf(err)

	err = conn.Ping(ctx)
	helper.PanicIf(err)

	return conn
}
