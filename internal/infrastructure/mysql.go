package infrastructure

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	client *sql.DB
}

func NewMySQLClient(user, password, dbname string) (*mysql, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", user, password, dbname)
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
