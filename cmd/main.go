package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rommomm123321/go-rest-api/internal/app"
	"github.com/rommomm123321/go-rest-api/internal/handler"
	"github.com/rommomm123321/go-rest-api/internal/repository"
	"github.com/rommomm123321/go-rest-api/internal/service"
	"github.com/rommomm123321/go-rest-api/pkg/config"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(cfg.DB.DSN)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	routes := handlers.InitRoutes()
	server := app.NewServer()
	if err := server.Run(cfg.Server.Port, routes); err != nil {
		log.Fatalf("Server error: %s", err.Error())
	}

}
