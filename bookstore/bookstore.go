package bookstore

import "errors"

type Book struct {
	Title  string
	Author string
	Copies int
	ID     int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func GetAllBooks(c []Book) []Book {
	return c
}

func GetBook(catalog []Book, i int) Book {
	for _, b := range catalog {
		if b.ID == i {
			return b
		}
	}
	return Book{}
}
