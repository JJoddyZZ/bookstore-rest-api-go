package infrastructure

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jzavala-globant/bookstore-rest-api-go/config"
)

type mysql struct {
	client *sql.DB
}

func NewMySQLClient(config *config.DB) (*mysql, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Password, config.Host, config.Port, config.Schema)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &mysql{
		db,
	}, nil
}

func (m *mysql) Ping() error {
	return m.client.Ping()
}
