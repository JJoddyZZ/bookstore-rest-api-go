package services

import (
	"context"

	"github.com/jzavala-globant/bookstore-rest-api-go/internal/models"
)

type Repositories interface {
	ListBooks(context.Context) (*models.Book, error)
}

type bookstore struct {
	r Repositories
}

func NewBookstoreService(r Repositories) *bookstore {
	return &bookstore{
		r,
	}
}

func (b *bookstore) ListBooks(context.Context) (*models.APIResponse, error) {
	return nil, nil
}
