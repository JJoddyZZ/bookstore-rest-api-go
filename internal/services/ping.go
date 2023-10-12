package services

import (
	"fmt"

	"github.com/jzavala-globant/bookstore-rest-api-go/internal/interfaces"
	"github.com/rs/zerolog"
)

type ping struct {
	log *zerolog.Logger
	r   interfaces.PingRepository
}

func NewPingService(r interfaces.PingRepository, log *zerolog.Logger) *ping {
	return &ping{
		log,
		r,
	}
}

func (p *ping) Ping() error {
	if err := p.r.Ping(); err != nil {
		return fmt.Errorf("error executing ping: %v", err)
	}
	return nil
}
