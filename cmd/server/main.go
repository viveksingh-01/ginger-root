package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/viveksingh-01/ginger-root/internal/config"
	"github.com/viveksingh-01/ginger-root/internal/database"
)

func main() {
	log.Println("Welcome to Ginger API.")
	cfg, _ := config.Load()

	database.Connect(cfg.MongoURI)

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.SetTrustedProxies(nil)

	log.Println("Server started on port:", cfg.Port)
	r.Run(":" + cfg.Port)
}
