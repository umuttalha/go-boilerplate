package main

import (
	"log"
	"net/http"

	"go-boilerplate/internal/config"
	"go-boilerplate/internal/database"
	"go-boilerplate/internal/handler"
	"go-boilerplate/internal/repository"
	"go-boilerplate/internal/service"
)

func main() {
	cfg := config.Load()

	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Run migrations
	if err := database.RunMigrations(db, "./migrations"); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	// Register routes
	mux := http.NewServeMux()
	mux.HandleFunc("/health", h.HealthCheck)

	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: mux,
	}

	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
