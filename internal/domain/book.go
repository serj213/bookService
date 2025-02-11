package domain

type Book struct {
	Id int
	Title string
	Author string
	CategoryId int
}


func NewBookDomain(id int, title string, author string, categoryId int) Book{
	return Book{
		Id: id,
		Title: title,
		Author: author,
		CategoryId: categoryId,
	}
}