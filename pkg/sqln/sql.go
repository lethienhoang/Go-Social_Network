package sqln

import (
	"database/sql"
)

type SQL struct {
	DB *sql.DB
}

func New(kind string, dsn string) (*SQL, error) {
	conn, err := sql.Open(kind, dsn)

	if err != nil {
		return nil, err
	}

	s := &SQL{
		DB: conn,
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	conn.SetConnMaxLifetime(_defaultConnTimeout)
	return s, nil
}
