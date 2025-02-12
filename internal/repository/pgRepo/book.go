package pgrepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
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

	err := pg.HandlePgTransaction(ctx, func(pgTX pgx.Tx) error {

		var categoryId int
		const querySelect = "SELECT id FROM categories WHERE id=$1"

		err := r.db.QueryRow(ctx, querySelect, category_id).Scan(&categoryId)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return fmt.Errorf("failed category_id=%d not found", category_id)
			}

			return fmt.Errorf("failed select category_id=%d: %w", category_id, err)
		}

		// insert new book
		const queryInsert = "INSERT INTO books(title, author, categories_id) VALUES ($1,$2,$3) RETURNING id,title,author,categories_id"
		err = r.db.QueryRow(ctx, queryInsert, title, author, categoryId).Scan(&newBook.Id, &newBook.Title, &newBook.Author, &newBook.CategoryId)
		if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == repository.PgCodeDublicate {
				return fmt.Errorf("%s: %w", pgErr.ColumnName, repository.ErrBookExists)
			}
		}
		return fmt.Errorf("failed create book: %w", err)
		}

		return nil

	}, r.db)
	
	if err != nil {
		return domain.Book{}, fmt.Errorf("failed transactions: %w", err)
	}
	return newBook, nil
}