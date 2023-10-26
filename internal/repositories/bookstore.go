package repositories

import (
	"context"
	"fmt"

	"github.com/jzavala-globant/bookstore-rest-api-go/internal/interfaces"
	"github.com/jzavala-globant/bookstore-rest-api-go/internal/models"
	"github.com/rs/zerolog"
)

const (
	booksTable = "books"
)

type bookstore struct {
	log *zerolog.Logger
	db  interfaces.DBClient
}

func NewBookstoreRepository(log *zerolog.Logger, db interfaces.DBClient) *bookstore {
	return &bookstore{
		log,
		db,
	}
}

func (b *bookstore) ListBooks(ctx context.Context) ([]models.Book, error) {
	var books []models.Book
	query := fmt.Sprintf("SELECT * FROM %s", booksTable)

	res, err := b.db.Query(query)
	if err != nil {
		return books, err
	}

	for res.Next() {
		var book models.Book
		err := res.Scan(
			&book.ID,
			&book.ISBN,
			&book.Title,
			&book.Author,
			&book.Genre,
			&book.Price,
			&book.Quantity,
		)
		if err != nil {
			return books, err
		}
		books = append(books, book)
	}

	return books, nil
}
