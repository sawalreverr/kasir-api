package main

import (
	"basic-go-api/config"
	"basic-go-api/internal/handler"
	"basic-go-api/internal/infra/database"
	"basic-go-api/internal/repository"
	"basic-go-api/internal/service"
	"log"
	"net/http"
)

func main() {
	// load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// init db
	db, err := database.NewPostgresDB(cfg.DBConn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// init repository
	categoryRepo := repository.NewCategoryRepository(db)

	// init service
	categoryService := service.NewCategoryService(categoryRepo)

	// init handler
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// setup router
	mux := http.NewServeMux()

	mux.HandleFunc("/categories", categoryHandler.Categories)
	mux.HandleFunc("/categories/", categoryHandler.CategoryByID)

	log.Printf("Server running on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
