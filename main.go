package main

import (
	"codewithumam-go-task1/database"
	"codewithumam-go-task1/handlers"
	"codewithumam-go-task1/internal/middleware"
	"codewithumam-go-task1/repositories"
	"codewithumam-go-task1/services"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}


func main() {
	router := http.NewServeMux()

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	// init db
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// product handler
	productRepo := repositories.NewProductRepository(db)
	productSvc := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productSvc)

	// category handler
	categoryRepo := repositories.NewCategoryRepository(db)
	categorySvc := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categorySvc)

	// root
	router.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))

	// Products
	router.Handle("GET /api/products", http.HandlerFunc(productHandler.GetAll))
	router.Handle("POST /api/products", http.HandlerFunc(productHandler.Create))
	router.Handle("GET /api/products/{id}", http.HandlerFunc(productHandler.GetByID))
	router.Handle("PUT /api/products/{id}", http.HandlerFunc(productHandler.Update))
	router.Handle("DELETE /api/products/{id}", http.HandlerFunc(productHandler.Delete))

	// Categories
	router.Handle("GET /api/categories", http.HandlerFunc(categoryHandler.GetCategories))
	router.Handle("GET /api/categories/{id}", http.HandlerFunc(categoryHandler.GetCategoryByID))
	router.Handle("POST /api/categories", http.HandlerFunc(categoryHandler.CreateCategory))
	router.Handle("PUT /api/categories/{id}", http.HandlerFunc(categoryHandler.UpdateCategory))
	router.Handle("DELETE /api/categories/{id}", http.HandlerFunc(categoryHandler.DeleteCategory))

	// Apply request logger middleware
	handler := middleware.RequestLogger(router)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Port),
		Handler: handler,
	}

	log.Printf("Server is running on port %s", config.Port)
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server: ", err)
		return
	}
}
