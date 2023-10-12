package repositories

import (
	"fmt"

	"github.com/rs/zerolog"
)

type ping struct {
	log *zerolog.Logger
	db  DBClient
}

func NewPingRepository(log *zerolog.Logger, db DBClient) *ping {
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
