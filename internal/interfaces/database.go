package interfaces

type DBClient interface {
	Ping() error
	Execute(statement string)
	Query(statement string) (DBRow, error)
	Close() error
}

type DBRow interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close() error
}
