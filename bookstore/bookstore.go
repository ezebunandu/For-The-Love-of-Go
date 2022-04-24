package bookstore

import (
	"errors"
	"fmt"
)

// Represents a book in the bookstore
type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
	category        string
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
func (catalog Catalog) GetAllBooks() Catalog {
	c := catalog
	return c
}

func (catalog Catalog) GetBook(ID int) (Book, error) {
	book, ok := catalog[ID]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}
	return book, nil
}

func (b Book) NetPriceCents() int {
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving
}

func (book *Book) SetPriceCents(newPrice int) error {
	if newPrice < 0 {
		return fmt.Errorf("bad price %d (must not be negative)", newPrice)
	}
	book.PriceCents = newPrice
	return nil
}

func (b *Book) SetCategory(category string) error {
	if category != "Autobiography" {
		return fmt.Errorf("unknown category %q", category)
	}
	b.category = category
	return nil
}

func (b Book) Category() string {
	return b.category
}
