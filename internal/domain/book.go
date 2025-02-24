package domain

import "time"

type Book struct {
	Id         int64
	Title      string
	Author     string
	CategoryId int
	UpdateAt time.Time
	CreateAt time.Time
}

func NewBookDomain(id int, title string, author string, categoryId int, updateAt time.Time, createAt time.Time) Book {
	return Book{
		Id:         int64(id),
		Title:      title,
		Author:     author,
		CategoryId: categoryId,
		UpdateAt: updateAt,
		CreateAt: createAt,
	}
}
