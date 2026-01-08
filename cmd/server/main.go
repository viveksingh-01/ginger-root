package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/viveksingh-01/ginger-root/internal/database"
)

func main() {
	log.Println("Welcome to Ginger API.")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect(os.Getenv("MONGODB_URI"))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.SetTrustedProxies(nil)

	log.Println("Server started on port:", port)
	r.Run(":" + port)
}
