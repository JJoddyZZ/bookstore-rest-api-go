package controllers

import (
	"net/http"

	"github.com/jzavala-globant/bookstore-rest-api-go/internal/interfaces"
	"github.com/rs/zerolog"
)

type pingController struct {
	log *zerolog.Logger
	s   interfaces.PingService
}

func NewPingController(s interfaces.PingService, log *zerolog.Logger) *pingController {
	return &pingController{
		log,
		s,
	}
}

func (b *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	err := b.s.Ping()
	if err != nil {
		b.log.Err(err).Msg("error sending ping signal")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
