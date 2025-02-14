package domain

type Book struct {
	Id         int64
	Title      string
	Author     string
	CategoryId int
}

func NewBookDomain(id int, title string, author string, categoryId int) Book {
	return Book{
		Id:         int64(id),
		Title:      title,
		Author:     author,
		CategoryId: categoryId,
	}
}
