package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mikolabarkouski/calculator/config"
	"github.com/mikolabarkouski/calculator/internal/api"
	"github.com/mikolabarkouski/calculator/internal/app"
	"github.com/mikolabarkouski/calculator/internal/repo"
)

func main() {
	cfg := config.LoadConfig()

	log.Printf("PACKAGES DEFAULT:%v", cfg.PackagesDefault)
	repository := repo.NewRepository(cfg.PackagesDefault)

	// Initialize application layer
	application := app.NewApp(repository)

	// Initialize API handlers
	handler := api.NewHandler(application)

	log.Printf("Starting server on :%s\n", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), handler.Router()); err != nil {
		log.Fatal("Server failed:", err)
	}
}
