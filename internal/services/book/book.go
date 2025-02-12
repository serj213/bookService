package book

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/serj213/bookService/internal/domain"
)


type BookRepository interface {
	Create(ctx context.Context, title string, author string, category_id int) (domain.Book, error)
}

type BookService struct {
	log *slog.Logger
	repo BookRepository
}

func NewBookService(log *slog.Logger, repo BookRepository) BookService{
	return BookService{
		log: log,
		repo: repo,
	}
}

func (s BookService) Create(ctx context.Context, title string, author string, categoryId int64) (domain.Book, error) {

	log := s.log.With(slog.String("op", "service.book.Create"))

	book, err := s.repo.Create(ctx, title, author, int(categoryId))

	if err != nil {
		log.Error("failed create book", err)
		return domain.Book{}, fmt.Errorf("failed create book: %w", err)
	}
	return book, nil
}

func (s BookService) Delete(ctx context.Context, id int)  error {
	return nil
}

func (s BookService) GetById(ctx context.Context, id int) (domain.Book, error) {
	return domain.Book{}, nil
}

func (s BookService) GetAllBooks(ctx context.Context) ([]domain.Book, error){
	return []domain.Book{}, nil
}

func (s BookService) Update(ctx context.Context, id int, categoryId int) (domain.Book, error){
	return domain.Book{}, nil
}