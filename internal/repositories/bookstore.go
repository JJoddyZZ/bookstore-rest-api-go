package repositories

import (
	"context"

	"github.com/jzavala-globant/bookstore-rest-api-go/internal/models"
	"github.com/rs/zerolog"
)

type dbClient interface {
}

type bookstore struct {
	log *zerolog.Logger
	db  dbClient
}

func NewBookstoreRepository(log *zerolog.Logger) *bookstore {
	return &bookstore{
		log: log,
		db:  nil,
	}
}

func (b *bookstore) ListBooks(ctx context.Context) (*models.Book, error) {
	return nil, nil
}
