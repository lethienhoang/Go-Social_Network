package postgres

import (
	"fmt"

	sqln "go-social-network.com/v1/pkg/sqln"
)

type Postgres struct {
	Postgresql *sqln.SQL
}

func New(url string) (*Postgres, error) {
	dsn := fmt.Sprintf("postgresql://%s", url)
	pg, err := sqln.New("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		Postgresql: pg,
	}, nil
}
