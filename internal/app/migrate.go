package app

import (
	"errors"
	"flag"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	_defaultAttempts = 20
	_defaultTimeout  = time.Second
)

func RunMigrate() {
	var (
		dbURL  string
		action string
	)

	fs := flag.NewFlagSet("flag", flag.ExitOnError)
	fs.StringVar(&dbURL, "dbURL", "", "database address connection string")
	fs.StringVar(&action, "action", "", "action to execute")

	// dbURL += "?sslmode=disable"

	var (
		attempts = _defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	if err = fs.Parse(os.Args[1:]); err != nil {
		log.Printf("parsing failed: %v", err)
	}

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

	if action == "up" {
		err = m.Up()
	} else if action == "down" {
		err = m.Down()
	}

	defer m.Close()

	if err != nil {
		log.Printf("migration failed: %v -> %v", action, err)
		return
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("migration no change")
		return
	}

	log.Printf("Migrate: %v succeeded", action)
}
