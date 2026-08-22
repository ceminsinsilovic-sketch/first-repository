package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connect(ctx context.Context)(*pgx.Conn, error){
	return pgx.Connect(ctx, "postgres://postgres:1234@localhost:5432/postgres")
}