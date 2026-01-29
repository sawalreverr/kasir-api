package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgresDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	// test connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// set connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Database connected successfully")
	return db, nil
}
