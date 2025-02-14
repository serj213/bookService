package book

import (
	"context"
	"log/slog"

	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/serj213/bookService/internal/domain"
)

type BookRepository interface {
	Create(ctx context.Context, title string, author string, category_id int) (domain.Book, error)
}

type BookService struct {
	log       *slog.Logger
	repo      BookRepository
	trManager *manager.Manager
}

func NewBookService(log *slog.Logger, repo BookRepository, trManager *manager.Manager) BookService {
	return BookService{
		log:       log,
		repo:      repo,
		trManager: trManager,
	}
}

func (s BookService) Create(ctx context.Context, title string, author string, categoryId int64) (domain.Book, error) {

	log := s.log.With(slog.String("op", "service.book.Create"))

	book, err := s.repo.Create(ctx, title, author, int(categoryId))
	if err != nil {
		log.Error("failed create book", "error", err)
		return domain.Book{}, err
	}
	return book, nil
}

func (s BookService) Delete(ctx context.Context, id int) error {
	return nil
}

func (s BookService) GetById(ctx context.Context, id int) (domain.Book, error) {
	return domain.Book{}, nil
}

func (s BookService) GetAllBooks(ctx context.Context) ([]domain.Book, error) {
	return []domain.Book{}, nil
}

func (s BookService) Update(ctx context.Context, id int, categoryId int) (domain.Book, error) {
	return domain.Book{}, nil
}
