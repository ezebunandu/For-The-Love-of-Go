package bookstore

// Represents a book in the bookstore
type Book struct {
	Title  string
	Author string
	Copies int
}

func Buy(b Book) Book {
	b.Copies--
	return b
}
