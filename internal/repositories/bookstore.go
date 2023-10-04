package repositories

import (
	"context"

	"github.com/jzavala-globant/bookstore-rest-api-go/internal/models"
)

type bookstore struct {
}

func NewBookstoreRepository() *bookstore {
	return &bookstore{}
}

func (b *bookstore) ListBooks(ctx context.Context) (*models.Book, error) {
	return nil, nil
}
