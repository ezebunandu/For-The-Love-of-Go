package bookstore

import (
	"errors"
	"fmt"
)

// Represents a book in the bookstore
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
func GetAllBooks(catalog map[int]Book) map[int]Book {
	c := catalog
	return c
}

func GetBook(catalog map[int]Book, ID int) (Book, error) {
	book, ok := catalog[ID]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}
	return book, nil
}
