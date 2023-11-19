package main

import (
	"fmt"
	"github.com/pressly/goose"
	"log"
	"net/http"
	"shopper-list-backend/config"
	"shopper-list-backend/db"
	"shopper-list-backend/router"
)

// main is the entry point of the application
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dbConnected, err := db.Connect(cfg.DatabaseConfig)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer dbConnected.Close()

	if err := goose.SetDialect("mysql"); err != nil {
		log.Fatal(err)
	}
	if err := goose.Up(db.DB, "migrations"); err != nil {
		log.Fatal(err)
	}

	r := router.InitializeRouter()

	fmt.Print("Server is running on port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
