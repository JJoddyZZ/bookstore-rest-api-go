package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jzavala-globant/bookstore-rest-api-go/internal/interfaces"
	"github.com/rs/zerolog"
)

type bookstoreController struct {
	log *zerolog.Logger
	s   interfaces.BookstoreService
}

func NewBookstoreController(s interfaces.BookstoreService, log *zerolog.Logger) *bookstoreController {
	return &bookstoreController{
		log,
		s,
	}
}

func (b *bookstoreController) ListBooks(w http.ResponseWriter, r *http.Request) {
	resp, err := b.s.ListBooks(r.Context())
	if err != nil {
		b.log.Err(err).Msg("error fetching books")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp.Body)
}
