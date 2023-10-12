package repositories

import (
	"context"

	"github.com/jzavala-globant/bookstore-rest-api-go/internal/models"
	"github.com/rs/zerolog"
)

type DBClient interface {
	Ping() error
}

type bookstore struct {
	log *zerolog.Logger
	db  DBClient
}

func NewBookstoreRepository(log *zerolog.Logger, db DBClient) *bookstore {
	return &bookstore{
		log,
		db,
	}
}

func (b *bookstore) ListBooks(ctx context.Context) (models.Book, error) {
	return models.Book{}, nil
}
