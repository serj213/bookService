package pgrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/serj213/bookService/internal/domain"
	"github.com/serj213/bookService/internal/repository"
	"github.com/serj213/bookService/internal/repository/model"
	grpcerror "github.com/serj213/bookService/pkg/grpcError"
	"github.com/serj213/bookService/pkg/pg"
)

type bookRepo struct {
	db *pg.PgDb
}

func NewBookRepo(db *pg.PgDb) bookRepo {
	return bookRepo{
		db: db,
	}
}

func (r bookRepo) Create(ctx context.Context, title string, author string, category_id int) (domain.Book, error) {
	var bookModel model.BookModel
	// insert new book
	const queryInsert = "INSERT INTO books(title, author, categories_id) SELECT $1,$2,$3 WHERE EXISTS (SELECT 1 FROM categories WHERE id = $3) RETURNING id, create_at"
	err := r.db.QueryRow(ctx, queryInsert, title, author, category_id).Scan(bookModel)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == repository.PgCodeDublicate {
				return domain.Book{}, grpcerror.ErrBookExists
			}
		}
		return domain.Book{}, fmt.Errorf("failed create book: %w", err)
	}

	newBook := bookToDomain(bookModel)

	return newBook, nil
}

func (r bookRepo) Delete(ctx context.Context, id int) error {
	const query = "DELETE FROM books WHERE id=$1"
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return grpcerror.ErrBookNotFound
	}
	return nil
}

func (r bookRepo) Update(ctx context.Context, book domain.Book) (domain.Book, error){

	update_at := time.Now()

	const query = `UPDATE books SET 
		title = COALESCE($2, title), author = COALESCE($3, author), 
		category_id = COALESCE($4, category_id), update_ad = COALESCE($5, update_at)
		WHERE id=$1
		RETURNING *
		`

	var bookModel model.BookModel

	err := r.db.Pool.QueryRow(ctx, query, book.Id, book.Title, book.Author, book.CategoryId, update_at).Scan(&bookModel)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Book{}, grpcerror.ErrBookNotFound
		}
		return domain.Book{}, fmt.Errorf("failed update book: %w", err)
	}
	updatedBook := bookToDomain(bookModel)
	
	return updatedBook, nil
}

func (r bookRepo) GetBookById(ctx context.Context, id int) (domain.Book, error) {
	var bookModel model.BookModel
	const query = `SELECT id, title, author, categories_id, created_at, updated_at FROM books WHERE id = $1`
	
	err := r.db.QueryRow(ctx, query, id).
		Scan(&bookModel.Id, &bookModel.Title, &bookModel.Author, &bookModel.CategoryId, &bookModel.CreateAt, &bookModel.UpdateAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Book{}, grpcerror.ErrBookNotFound
		}
		return domain.Book{}, fmt.Errorf("failed get book by id: %w", err)
	}

	book := bookToDomain(bookModel)

	return book, nil
}