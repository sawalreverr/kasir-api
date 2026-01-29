package main

import (
	"basic-go-api/config"
	"basic-go-api/internal/infra/database"
	"log"
)

func main() {
	// load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// init db
	_, err = database.NewPostgresDB(cfg.DBConn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
}
