package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/viveksingh-01/ginger-root/internal/config"
	"github.com/viveksingh-01/ginger-root/internal/database"
)

func main() {
	log.Println("Welcome to Ginger API.")

	// Load configurations
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err.Error())
	}

	// Connect to DB
	database.Connect(cfg.MongoURI)

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.SetTrustedProxies(nil)

	log.Println("Server started on port:", cfg.Port)
	r.Run(":" + cfg.Port)
}
