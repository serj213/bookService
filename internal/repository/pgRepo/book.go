package pgrepo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/serj213/bookService/internal/domain"
	"github.com/serj213/bookService/internal/repository"
	"github.com/serj213/bookService/pkg/pg"
)


type bookRepo struct {
	db *pg.PgDb
}

func NewBookRepo(db *pg.PgDb) bookRepo{
	return bookRepo{
		db: db,
	}
}


func (r bookRepo) Create(ctx context.Context, title string, author string, category_id int) (domain.Book, error) {
	var newBook domain.Book
	const query = "INSERT INTO books(title, author, category_id) VALUES ($1,$2,$3) RETURNING id,title,category_id"

	err := r.db.QueryRow(ctx, query, title, author, category_id).Scan(&newBook)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == repository.PgCodeDublicate {
				return domain.Book{}, fmt.Errorf("%s: %w", pgErr.ColumnName, repository.ErrBookExists)
			}
		}

		return domain.Book{}, fmt.Errorf("failed create book: %w", err)
	}

	return newBook, nil
}