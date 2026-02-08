package main

import (
	"basic-go-api/config"
	"basic-go-api/internal/handler"
	"basic-go-api/internal/infra/database"
	"basic-go-api/internal/infra/seed"
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
	defer db.Close()

	// run seed
	if cfg.Env == "development" && cfg.Seed {
		seed.Run(db)
	}

	// init repository
	categoryRepo := repository.NewCategoryRepository(db)
	productRepo := repository.NewProductRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	// init service
	categoryService := service.NewCategoryService(categoryRepo)
	productService := service.NewProductService(productRepo)
	transactionService := service.NewTransactionService(transactionRepo, productRepo, db)

	// init handler
	categoryHandler := handler.NewCategoryHandler(categoryService)
	productHandler := handler.NewProductHandler(productService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	// setup router
	mux := http.NewServeMux()

	mux.HandleFunc("/categories", categoryHandler.Categories)
	mux.HandleFunc("/categories/", categoryHandler.CategoryByID)
	mux.HandleFunc("/products", productHandler.Products)
	mux.HandleFunc("/products/", productHandler.ProductByID)
	mux.HandleFunc("/checkout", transactionHandler.Transactions)

	log.Printf("Server running on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
