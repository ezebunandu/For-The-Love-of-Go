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

// a catalog contains books mapped to their IDs
type Catalog map[int]Book

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}
func GetAllBooks(catalog Catalog) Catalog {
	c := catalog
	return c
}

func GetBook(catalog Catalog, ID int) (Book, error) {
	book, ok := catalog[ID]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}
	return book, nil
}
