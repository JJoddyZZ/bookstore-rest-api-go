package services

import (
	"context"

	"github.com/jzavala-globant/bookstore-rest-api-go/internal/interfaces"
	"github.com/jzavala-globant/bookstore-rest-api-go/internal/models"
	"github.com/rs/zerolog"
)

type bookstore struct {
	log *zerolog.Logger
	r   interfaces.BookstoreRepository
}

func NewBookstoreService(r interfaces.BookstoreRepository, log *zerolog.Logger) *bookstore {
	return &bookstore{
		log,
		r,
	}
}

func (b *bookstore) ListBooks(context.Context) (models.APIResponse, error) {
	return models.APIResponse{}, nil
}
