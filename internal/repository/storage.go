package repository

import "errors"


var (
	ErrBookExists = errors.New("book is exists")
	ErrBookNotFound = errors.New("book not found")
)

const(
	PgCodeDublicate = "23505"
)