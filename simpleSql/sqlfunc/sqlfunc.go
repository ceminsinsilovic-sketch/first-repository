package sqlfunc

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	SQLQuery := `
	CREATE TABLE IF NOT EXISTS books(
		id SERIAL PRIMARY KEY,
		name VARCHAR(70) NOT NULL,
		author VARCHAR(100) NOT NULL,
		review VARCHAR(1000) NOT NULL,
		year INT NOT NULL,
		is_read BOOLEAN NOT NULL,
		is_there TIMESTAMP NOT NULL,
		full_read TIMESTAMP
		);
	`
	_, err := conn.Exec(ctx, SQLQuery)
	return err 
}
