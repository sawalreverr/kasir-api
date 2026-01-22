package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// in memory storage
var CATEGORIES = []Category{
	{ID: 1, Name: "food", Description: "Foods"},
	{ID: 2, Name: "electronics", Description: "Electronics"},
	{ID: 3, Name: "book", Description: "Books"},
}

// response
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func main() {
	// GET /categories -> get all categories
	// POST /categories -> add new category
	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case "GET":
			json.NewEncoder(w).Encode(Response{
				Success: true,
				Message: "all categories",
				Data:    CATEGORIES,
			})
		case "POST":
			var newCategory Category
			if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(Response{
					Success: false,
					Message: err.Error(),
					Data:    nil,
				})
				return
			}

			newCategory.ID = len(CATEGORIES) + 1
			CATEGORIES = append(CATEGORIES, newCategory)

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(Response{
				Success: true,
				Message: "category created",
				Data:    newCategory,
			})
		}
	})

	// GET /categories/{id} -> get category by id
	// PUT /categories/{id} -> update category by id
	// DELETE /categories/{id} -> delete category by id
	http.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		idParam := strings.TrimPrefix(r.URL.Path, "/categories/")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})
			return
		}

		switch r.Method {
		case "GET":
			for _, category := range CATEGORIES {
				if category.ID == id {
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(Response{
						Success: true,
						Message: "category found",
						Data:    category,
					})
					return
				}
			}

			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Response{
				Success: false,
				Message: "category not found",
				Data:    nil,
			})
			return
		case "PUT":
			var newCategory Category
			if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(Response{
					Success: false,
					Message: err.Error(),
					Data:    nil,
				})
				return
			}

			for idx := range CATEGORIES {
				if CATEGORIES[idx].ID == id {
					CATEGORIES[idx] = newCategory

					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(Response{
						Success: true,
						Message: "category updated",
						Data:    newCategory,
					})
					return
				}
			}

			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Response{
				Success: false,
				Message: "category not found",
				Data:    nil,
			})
			return
		case "DELETE":
			for idx, category := range CATEGORIES {
				if category.ID == id {
					CATEGORIES = append(CATEGORIES[:idx], CATEGORIES[idx+1:]...)

					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(Response{
						Success: true,
						Message: "category deleted",
						Data:    category,
					})
					return
				}
			}

			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Response{
				Success: false,
				Message: "category not found",
				Data:    nil,
			})
			return
		}
	})

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed starting server %v", err)
	}
}
