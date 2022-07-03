package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go-social-network.com/v1/web"
)

func main() {
	run()
}

func run() error {

	var (
		addr string
		// sqlAddr    string
		// sessionKey string
	)

	fs := flag.NewFlagSet("flag", flag.ExitOnError)
	// fs.StringVar(&sessionKey, "session-key")
	// fs.StringVar(&sqlAddr, "sqlAddr")
	fs.StringVar(&addr, "addr", ":4000", "Https server address")

	if err := fs.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("error parsing: %w", err)
	}

	logger := log.New(os.Stderr, "", log.Lshortfile|log.Ldate|log.Ltime)

	webHanlder := &web.Handler{
		Logger: logger,
	}

	srv := &http.Server{
		Handler:      webHanlder.ServeHTTP(),
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	defer srv.Close()

	fmt.Printf("Server listening on %s\n", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and serve: %w", err)
	}

	return nil
}
