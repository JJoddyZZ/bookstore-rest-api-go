package services

import (
	"context"

	"github.com/jzavala-globant/bookstore-rest-api-go/internal/models"
	"github.com/rs/zerolog"
)

type Repositories interface {
	ListBooks(context.Context) (*models.Book, error)
}

type bookstore struct {
	log *zerolog.Logger
	r   Repositories
}

func NewBookstoreService(r Repositories, log *zerolog.Logger) *bookstore {
	return &bookstore{
		log,
		r,
	}
}

func (b *bookstore) ListBooks(context.Context) (*models.APIResponse, error) {
	return nil, nil
}
