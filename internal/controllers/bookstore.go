package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jzavala-globant/bookstore-rest-api-go/internal/models"
)

type Services interface {
	ListBooks(context.Context) (*models.APIResponse, error)
}

type bookstore struct {
	s Services
}

func NewBookstoreController(s Services) *bookstore {
	return &bookstore{
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
