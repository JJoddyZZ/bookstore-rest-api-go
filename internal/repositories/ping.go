package repositories

import (
	"fmt"

	"github.com/jzavala-globant/bookstore-rest-api-go/internal/interfaces"
	"github.com/rs/zerolog"
)

type ping struct {
	log *zerolog.Logger
	db  interfaces.DBClient
}

func NewPingRepository(log *zerolog.Logger, db interfaces.DBClient) *ping {
	return &ping{
		log,
		db,
	}
}

func (p *ping) Ping() error {
	if err := p.db.Ping(); err != nil {
		return fmt.Errorf("error sending ping to db: %v", err)
	}
	return nil
}
