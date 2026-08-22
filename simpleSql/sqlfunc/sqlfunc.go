package sqlfunc

import (
	"context"
	"fmt"

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

func InsertBook(ctx context.Context, conn *pgx.Conn, book Book) error {
	SQLQuery := `
	INSERT INTO books(name, author, review, year, is_read, is_there, full_read)
	VALUES($1, $2, $3, $4, $5, $6, $7);
	`
	_, err := conn.Exec(ctx,
		SQLQuery,
		book.Name,
		book.Author,
		book.Review,
		book.Year,
		book.IsRead,
		book.IsThere,
		book.FullRead,
	)
	return err
}

func UpdateData(ctx context.Context, conn *pgx.Conn, book Book) error {
	SQLQuery := `
	UPDATE books
	SET name=$1, author=$2, review=$3, year=$4, is_read=$5, is_there=$6, full_read=$7
	WHERE id=$8;
	`
	_, err := conn.Exec(
		ctx,
		SQLQuery,
		book.Name,
		book.Author,
		book.Review,
		book.Year,
		book.IsRead,
		book.IsThere,
		book.FullRead,
		book.ID,
	)
	return err
}

func Delete(ctx context.Context, conn *pgx.Conn, IDs []int) error {
	SQLQuery := `
	DELETE FROM books
	WHERE id = ANY($1)
	`
	_, err := conn.Exec(ctx, SQLQuery, IDs)
	return err
}

func OZON(ctx context.Context, conn *pgx.Conn) ([]Book, error) {
	SQLQuery := `
	SELECT id, name, author, review, year, is_read, is_there, full_read
	FROM books
	`
	rows, err := conn.Query(ctx, SQLQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	boooks := make([]Book, 0)
	for rows.Next() {
		var b Book
		if err := rows.Scan(
			&b.ID,
			&b.Name,
			&b.Author,
			&b.Review,
			&b.Year,
			&b.IsRead,
			&b.IsThere,
			&b.FullRead,
		); err != nil {
			return nil, err
		}
		boooks = append(boooks, b)
	}
	return boooks, nil
}

func Updatebook(ctx context.Context, conn *pgx.Conn, book Book) error {
	SQLQuery := `
	UPDATE books
	SET name=$1, author=$2, review=$3, year=$4, is_read=$5, is_there=$6, full_read=$7
	WHERE id=$8;
	`
	_, err := conn.Exec(ctx,
		SQLQuery,
		book.Name,
		book.Author,
		book.Review,
		book.Year,
		book.IsRead,
		book.IsThere,
		book.FullRead,
		book.ID,
	)
	return err

}
func ListPages(n int, ctx context.Context, conn *pgx.Conn)([]Book, error)  {
	offset := 0
	limit := 5
	SQLQuery := `
	SELECT * FROM books
	ORDER BY id ASC LIMIT $1
	OFFSET $2
	`
	books := make([]Book, 0)
	for i := 0; i < n; i++ {
		var b Book
		rows, err := conn.Query(ctx, SQLQuery, limit, offset)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next(){
			if err := rows.Scan(
				&b.ID,
				&b.Name,
				&b.Author,
				&b.Review,
				&b.Year,
				&b.IsRead,
				&b.IsThere,
				&b.FullRead,
			); err != nil{
				return nil, err
			}
		}
		books = append(books, b)
		
		fmt.Println(n, offset, limit, books)
		offset += limit
	}
	return books, nil
}
