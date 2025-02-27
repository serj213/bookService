package book

import (
	"github.com/serj213/bookService/internal/domain"
	bsv1 "github.com/serj213/bookService/pb/grpc/grpc"
)
	

func domainToGrpcFormat(book domain.Book) *bsv1.BookResponse{
	return &bsv1.BookResponse{
		Id: book.Id,
		Title: book.Title,
		Author: book.Author,
		CategoryId: int64(book.CategoryId),
	}
}