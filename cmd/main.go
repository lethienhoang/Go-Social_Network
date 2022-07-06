package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"go-social-network.com/v1/internal/db"
	service "go-social-network.com/v1/internal/services"
	"go-social-network.com/v1/web"
)

func main() {
	run()
}

func run() error {

	var (
		addr        string
		sqlAddr     string
		sqlPassword string
		// sessionKey string
	)

	fs := flag.NewFlagSet("flag", flag.ExitOnError)
	// fs.StringVar(&sessionKey, "session-key")
	fs.StringVar(&sqlAddr, "sqlAddr", "", "postgres address database")
	fs.StringVar(&sqlPassword, "sqlPass", "", "sql password")
	fs.StringVar(&addr, "addr", ":4000", "Https server address")

	var err error

	if err = fs.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("error parsing: %w", err)
	}

	dsn := fmt.Sprintf("postgresql://go-user:%s@%s", sqlPassword, sqlAddr)
	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	queries := db.New(dbConn)
	svc := &service.Service{
		Queries: queries,
	}

	logger := log.New(os.Stderr, "", log.Lshortfile|log.Ldate|log.Ltime)

	webHanlder := &web.Handler{
		Logger:  logger,
		Service: svc,
	}

	srv := &http.Server{
		Handler:      webHanlder.ServeHTTP(),
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	defer srv.Close()

	fmt.Printf("Server listening on %s\n", srv.Addr)
	err = srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and serve: %w", err)
	}

	return nil
}
