package simpleconnect

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connect(ctx context.Context) (*pgx.Conn, error) {
	connStr := "postgres://username:password@localhost:5432/databasename"
	return pgx.Connect(ctx, connStr)
}
