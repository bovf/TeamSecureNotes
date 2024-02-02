package main

import (
	"log"
	"teams-secure-notes/internal/handler"
	"teams-secure-notes/internal/repository"
	"teams-secure-notes/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize the Database
	repository.InitializeDB(cfg.MongoDBUser, cfg.MongoDBPassword, cfg.MongoDBHost, cfg.MongoDBPort, cfg.MongoDBName)

	// Setup Fiber app
	app := fiber.New()

	// Register routes
	app.Get("/messages", handler.GetMessages)
	app.Post("/messages", handler.CreateMessage)

	// Start server
	log.Fatal(app.Listen(":" + cfg.Port))
}

