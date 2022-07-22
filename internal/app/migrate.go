package app

import (
	"errors"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	_defaultAttempts = 20
	_defaultTimeout  = time.Second
)

func Execute(dbURL string, migrateAction string) {
	// dbURL += "?sslmode=disable"

	var (
		attempts = _defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New("file://internal/sqlc/migrations", dbURL)
		if err == nil {
			break
		}

		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		log.Printf("migration failed: %v", err)
		return
	}

	if migrateAction == "up" {
		err = m.Up()
	} else if migrateAction == "down" {
		err = m.Down()
	}

	defer m.Close()

	if err != nil {
		log.Printf("migration failed: %v -> %v", migrateAction, err)
		return
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("migration no change")
		return
	}

	log.Printf("Migrate: %v succeeded", migrateAction)
}
