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

func (s BookService) Create(ctx context.Context, title string, author string, category_id int) (domain.Book, error) {

	log := s.log.With(slog.String("op", "service.book.Create"))

	book, err := s.repo.Create(ctx, title, author, category_id)

	if err != nil {
		log.Error("failed create book", err)
		return domain.Book{}, fmt.Errorf("failed create book: %w", err)
	}
	return book, nil
}