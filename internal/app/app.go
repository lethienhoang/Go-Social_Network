package app

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

	"go-social-network.com/v1/internal/db"
	"go-social-network.com/v1/internal/services"
	"go-social-network.com/v1/pkg/postgres"
	"go-social-network.com/v1/web"
)

func Run() {
	var (
		dbURL   string
		address string
	)

	fs := flag.NewFlagSet("flag", flag.ExitOnError)
	// fs.StringVar(&sessionKey, "session-key")
	fs.StringVar(&dbURL, "dbURL", "", "postgres address database")
	fs.StringVar(&address, "address", ":4000", "Https server address")

	logger := log.New(os.Stderr, "", log.Lshortfile|log.Ldate|log.Ltime)

	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Printf("parsing failed: %v", err)
	}

	databaseConn, err := postgres.New(dbURL)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	queries := db.New(databaseConn.Postgresql.DB)
	svc := &services.Service{
		Queries: queries,
	}

	webHanlder := &web.Handler{
		Logger:  logger,
		Service: svc,
	}

	srv := &http.Server{
		Handler:      webHanlder.ServeHTTP(),
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	defer srv.Close()

	fmt.Printf("Server listening on %s\n", srv.Addr)
	err = srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Errorf("listen and serve: %w", err)
	}
}
