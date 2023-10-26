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

func (b *bookstore) ListBooks(ctx context.Context) (models.APIResponse, error) {
	var res models.APIResponse

	books, err := b.r.ListBooks(ctx)
	if err != nil {
		return res, err
	}

	res.Body = books
	return res, nil
}
