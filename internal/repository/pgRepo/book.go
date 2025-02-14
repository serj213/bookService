package pgrepo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/serj213/bookService/internal/domain"
	"github.com/serj213/bookService/internal/repository"
	grpcerror "github.com/serj213/bookService/pkg/grpcError"
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
	var bookId int
	// insert new book
	const queryInsert = "INSERT INTO books(title, author, categories_id) SELECT $1,$2,$3 WHERE EXISTS (SELECT 1 FROM categories WHERE id = $3) RETURNING id"
	err := r.db.QueryRow(ctx, queryInsert, title, author, category_id).Scan(&bookId)
	if err != nil {
	if pgErr, ok := err.(*pgconn.PgError); ok {
		if pgErr.Code == repository.PgCodeDublicate {
			return domain.Book{}, grpcerror.ErrBookExists
		}
	}
	return domain.Book{}, fmt.Errorf("failed create book: %w", err)
	}

	newBook := domain.NewBookDomain(bookId, title, author, category_id)

	return newBook, nil
}