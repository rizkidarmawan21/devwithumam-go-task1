package main

import (
	"codewithumam-go-task1/internal/category"
	"codewithumam-go-task1/internal/middleware"
	"fmt"
	"log"
	"net/http"
)

const port = 8083

func main() {
	router := http.NewServeMux()

	// category handler
	categoryHandler := category.NewHTTPHandler()

	// root
	router.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))

	router.Handle("GET /categories", http.HandlerFunc(categoryHandler.GetCategories))
	router.Handle("GET /categories/{id}", http.HandlerFunc(categoryHandler.GetCategoryByID))
	router.Handle("POST /categories", http.HandlerFunc(categoryHandler.CreateCategory))
	router.Handle("PUT /categories/{id}", http.HandlerFunc(categoryHandler.UpdateCategory))
	router.Handle("DELETE /categories/{id}", http.HandlerFunc(categoryHandler.DeleteCategory))

	// Apply request logger middleware
	handler := middleware.RequestLogger(router)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}

	log.Printf("Server is running on port %d", port)
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server: ", err)
		return
	}
}
