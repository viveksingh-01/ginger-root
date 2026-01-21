package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/viveksingh-01/ginger-root/internal/config"
	"github.com/viveksingh-01/ginger-root/internal/database"
	"github.com/viveksingh-01/ginger-root/internal/server"
)

func main() {
	log.Println("Welcome to Ginger API.")
	gin.SetMode(gin.ReleaseMode)

	// Load configurations
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err.Error())
	}

	// Connect to DB
	mongoClient, err := database.Connect(cfg.MongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err.Error())
	}
	db := mongoClient.Database(cfg.Database)

	// Create a new server and start it
	server := server.New(cfg, db)
	server.Start()
}
