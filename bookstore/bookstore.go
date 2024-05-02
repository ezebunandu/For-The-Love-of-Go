package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
	category        string
}

type Catalog map[int]Book

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func (c Catalog) GetAllBooks() []Book {
	result := []Book{}
	for _, b := range c {
		result = append(result, b)
	}
	return result
}

func (c Catalog) GetBook(ID int) (Book, error) {
	book, ok := c[ID]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}
	return book, nil
}

func (b Book) NetPriceCents() int {
	return b.PriceCents * (100 - b.DiscountPercent) / 100
}

func (b *Book) SetPriceCents(price int) error {
	if price < 0 {
		return errors.New("price cannot be negative")
	}
	b.PriceCents = price
	return nil
}

func (b *Book) SetCategory(category string) error {
	if category != "Autobiography" {
		return fmt.Errorf("unknown category: %s", category)
	}
	b.category = category
	return nil
}

func (b Book) Category() string {
	return b.category
}
