package infrastructure

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jzavala-globant/bookstore-rest-api-go/config"
	"github.com/jzavala-globant/bookstore-rest-api-go/internal/interfaces"
)

type mysql struct {
	client *sql.DB
}

type row struct {
	rows *sql.Rows
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

func (m *mysql) Query(query string) (interfaces.DBRow, error) {
	rows, err := m.client.Query(query)
	if err != nil {
		return nil, err
	}

	// defer rows.Close()

	res := &row{
		rows: rows,
	}

	return res, nil
}

func (m *mysql) Execute(statement string) {
	m.client.Exec(statement)
}

func (r *row) Scan(dest ...interface{}) error {
	err := r.rows.Scan(dest...)
	if err != nil {
		return err
	}
	return nil
}

func (r *row) Next() bool {
	return r.rows.Next()
}
