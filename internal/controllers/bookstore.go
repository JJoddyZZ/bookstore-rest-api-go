package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jzavala-globant/bookstore-rest-api-go/internal/models"
	"github.com/rs/zerolog"
)

type Services interface {
	ListBooks(context.Context) (*models.APIResponse, error)
}

type bookstore struct {
	log *zerolog.Logger
	s   Services
}

func NewBookstoreController(s Services, log *zerolog.Logger) *bookstore {
	return &bookstore{
		log,
		s,
	}
}

func (b *bookstore) ListBooks(w http.ResponseWriter, r *http.Request) {
	resp, err := b.s.ListBooks(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
