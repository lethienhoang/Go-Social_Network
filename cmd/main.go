package main

import (
	"flag"
	"log"
	"os"

	"go-social-network.com/v1/internal/app"
)

func main() {

	var (
		action        string
		dbURL         string
		address       string
		migrateAction string
	)

	fs := flag.NewFlagSet("flag", flag.ExitOnError)
	fs.StringVar(&action, "action", "", "action to execute")
	fs.StringVar(&dbURL, "dbURL", "", "postgres address database")
	fs.StringVar(&address, "address", ":4000", "Https server address")
	fs.StringVar(&migrateAction, "migrateAction", "", "action to execute")

	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Printf("parsing failed: %v", err)
	}

	if action == "migrate" {
		app.Execute(dbURL, migrateAction)
		return
	} else {
		app.Run(dbURL, address)
	}
}
